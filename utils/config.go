package utils

type Config struct {
	GuildID            string
	MongoConnectionURI string
	DiscordToken       string
	BoosterRoleID      string
	AdminRoleID        string
	ModeratorRoleID    string
}

func LoadConfig() (conf Config, err error) {
	conf = Config{}

	conf.GuildID, err = GetEnviromentVariable("codify.GuildID")

	if err != nil {
		return conf, err
	}

	conf.AdminRoleID, err = GetEnviromentVariable("codify.AdminRoleID")

	if err != nil {
		return conf, err
	}

	conf.ModeratorRoleID, err = GetEnviromentVariable("codify.ModeratorRoleID")

	if err != nil {
		return conf, err
	}

	conf.BoosterRoleID, err = GetEnviromentVariable("codify.BoosterRoleID")

	if err != nil {
		return conf, err
	}

	conf.DiscordToken, err = GetEnviromentVariable("discord.Token")

	if err != nil {
		return conf, err
	}

	conf.MongoConnectionURI, err = GetEnviromentVariable("mongo.Uri")

	if err != nil {
		return conf, err
	}

	return conf, err
}
