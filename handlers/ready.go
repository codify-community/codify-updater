package handlers

import (
	"context"
	f "fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/codify-community/internals/codify-updater/database"
	"github.com/codify-community/internals/codify-updater/log"
	"github.com/codify-community/internals/codify-updater/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func update(sess *discordgo.Session, db *mongo.Collection, conf *utils.Config) error {
	general := database.General{}
	guild, err := sess.State.Guild(conf.GuildID)

	log.Info(f.Sprintf("found main guild as: %s (%v)", guild.Name, guild.ID))

	if err != nil {
		return f.Errorf("failed to get main guild: %v", err)
	}

	members, err := utils.FetchAllMembers(sess, guild)
	if err != nil {
		return f.Errorf("failed to get members: %v", err)
	}

	log.Info(f.Sprintf("member_count: %v", guild.MemberCount))

	for _, member := range members {
		if role := getRole(conf, member); role != "" {
			entity := database.Entity{
				Role:   role,
				Member: database.FromDiscordMember(member),
			}

			general.SpecialMembers = append(general.SpecialMembers, entity)
		}
	}

	general.ChannelCount = len(guild.Channels)
	general.MemberCount = guild.MemberCount

	_, err = db.UpdateByID(context.Background(), "0", bson.D{bson.E{
		Key: "$set", Value: general,
	}})

	return err
}

func OnReady(root *mongo.Client, conf *utils.Config, session *discordgo.Session, ev *discordgo.Ready) {
	log.Info("ready")
	stats := root.Database("stats")
	log.Wait("pinging db...")
	utils.Try(stats.Client().Ping(context.Background(), readpref.Primary()))
	log.Info("database is ok!")
	general := stats.Collection("general")

	for {
		log.Wait("updating...")
		err := update(session, general, conf)
		if err != nil {
			log.Error(err.Error())
		} else {
			log.Info("updated.")
		}
		time.Sleep(time.Minute * 15)
	}
}
