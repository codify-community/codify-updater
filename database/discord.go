package database

import (
	"github.com/bwmarrin/discordgo"
	"github.com/codify-community/internals/codify-updater/utils"
)

func FromDiscordMember(member *discordgo.Member) Member {
	return Member{
		Id:        member.User.ID,
		Name:      utils.StringOr(member.Nick, member.User.Username),
		AvatarURL: member.AvatarURL(""),
	}
}
