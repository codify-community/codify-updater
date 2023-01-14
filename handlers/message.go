package handlers

import (
	"context"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/codify-community/internals/codify-updater/database"
	"github.com/codify-community/internals/codify-updater/log"
	"github.com/codify-community/internals/codify-updater/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func prepare(info *mongo.Collection, memberDoc bson.D, conf *utils.Config, authorID string, ev *discordgo.MessageCreate) *database.InfoEntity {
	var memberEntity database.InfoEntity

	err := info.FindOne(context.Background(), memberDoc).Decode(&memberEntity)

	if role := getRole(conf, ev.Member); role != "" {
		if err == mongo.ErrNoDocuments {
			_, err := info.InsertOne(context.Background(), database.InfoEntity{MemberID: authorID})
			utils.Try(err)
		}

		return &memberEntity
	} else {
		if err != nil {
			info.DeleteOne(context.Background(), memberDoc)
		}
	}

	return nil
}

func OnMessage(root *mongo.Client, conf *utils.Config, session *discordgo.Session, ev *discordgo.MessageCreate) {
	props := root.Database("props-v2")
	info := props.Collection("info")
	authorID := utils.StringOr(ev.Author.ID, ev.Message.Author.ID)
	memberDoc := bson.D{bson.E{Key: "member_id", Value: authorID}}

	if authorID == "" {
		log.Error("User is empty, something has happend. Cache not ready?")
		return
	}

	if entity := prepare(info, memberDoc, conf, authorID, ev); entity != nil {
		res := strings.Split(ev.Message.Content, " ")

		if len(res) < 2 {
			return
		}

		command := res[0]
		kind := res[1]
		rest := res[2:]

		switch command {
		case ".editar":
			fallthrough
		case ".edit":
			switch kind {
			case "bio":
				entity.Bio = strings.Join(rest, " ")
			case "github":
				entity.Github = strings.Join(rest, " ")
			case "occupation":
				entity.Occupation = strings.Join(rest, " ")
			case "skills":
				entity.Skills = rest
			default:
				session.MessageReactionAdd(ev.Message.ChannelID, ev.Message.ID, "❌")
				return
			}

			info.UpdateOne(context.Background(), memberDoc, bson.D{bson.E{Key: "$set", Value: entity}})
			session.MessageReactionAdd(ev.Message.ChannelID, ev.Message.ID, "👌")
		}
	}
}
