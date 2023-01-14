package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/codify-community/internals/codify-updater/database"
	"github.com/codify-community/internals/codify-updater/utils"
)

func getRole(conf *utils.Config, m *discordgo.Member) database.Role {
	if _, ok := utils.Has(m.Roles, conf.AdminRoleID); ok {
		return database.Admin
	}

	if _, ok := utils.Has(m.Roles, conf.ModeratorRoleID); ok {
		return database.Moderator
	}

	if _, ok := utils.Has(m.Roles, conf.BoosterRoleID); ok {
		return database.Booster
	}

	return ""
}
