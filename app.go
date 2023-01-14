package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/codify-community/internals/codify-updater/handlers"
	"github.com/codify-community/internals/codify-updater/log"
	"github.com/codify-community/internals/codify-updater/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load()
	conf, err := utils.LoadConfig()
	utils.Try(err)

	log.Wait("connecting to database ...")
	mongo, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.MongoConnectionURI))
	utils.Try(err)

	discord, err := discordgo.New("Bot " + conf.DiscordToken)
	utils.Try(err)

	discord.AddHandlerOnce(func(s *discordgo.Session, r *discordgo.Ready) {
		go handlers.OnReady(mongo, &conf, s, r)
	})

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Message != nil && m.Author != nil && m.Member != nil {
			handlers.OnMessage(mongo, &conf, s, m)
		}
	})

	utils.Try(err)

	defer discord.Close()

	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	utils.Try(discord.Open())

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
