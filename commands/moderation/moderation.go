package moderation

import (
	"github.com/RhykerWells/asbwig/common"
	"github.com/RhykerWells/asbwig/common/dcommand"
)

//go:generate sqlboiler --no-hooks psql

func ModerationSetup(cmdHandler *dcommand.CommandHandler) {
	common.InitSchema("Moderation", GuildModerationSchema...)
	cmdHandler.RegisterCommands(
	)
}

func GuildModerationAdd(guild_id string) {
	const query = `SELECT guild_id FROM moderation_config WHERE guild_id=$1`
	err := common.PQ.QueryRow(query, guild_id)
	if err != nil {
		guildModerationDefault(guild_id)
	}
}

func guildModerationDefault(guild_id string) {
	const query = `INSERT INTO moderation_config (guild_id) VALUES ($1)`
	common.PQ.Exec(query, guild_id)
}