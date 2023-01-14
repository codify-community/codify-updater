package utils

import (
	"github.com/bwmarrin/discordgo"
)

func FetchAllMembers(s *discordgo.Session, guild *discordgo.Guild) ([]*discordgo.Member, error) {
	members := []*discordgo.Member{}
	lastMember := ""

	for feched := 0; feched < guild.MemberCount; feched += 1000 {
		newMembers, err := s.GuildMembers(guild.ID, lastMember, 1000)

		if err != nil {
			return members, err
		}

		lastMember = newMembers[len(newMembers)-1].User.ID
		members = append(members, newMembers...)
	}

	return members, nil
}
