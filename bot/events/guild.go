package events

import (
	"context"

	"github.com/RhykerWells/asbwig/bot/prefix"
	"github.com/RhykerWells/asbwig/commands/economy"
	"github.com/RhykerWells/asbwig/common"
	"github.com/RhykerWells/asbwig/common/models"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// guildJoin is called when the bot is added to a new guild
// This adds the guild to the relevant database tables
func guildJoin(s *discordgo.Session, g *discordgo.GuildCreate) {
	log.WithFields(log.Fields{
		"guild":       g.ID,
		"owner":       g.OwnerID,
		"membercount": g.MemberCount,
	}).Infoln("Joined guild: ", g.Name)
	banned := isGuildBanned(g.ID)
	if banned {
		common.Session.GuildLeave(g.ID)
		return
	}
	prefix.GuildPrefix(g.ID)
	economy.GuildEconomyAdd(g.ID)
}

// guildLeave is called when the bot is removed from a guild
// This removes the guild from any tables that it is part of
func guildLeave(s *discordgo.Session, g *discordgo.GuildDelete) {
	if g.Unavailable {
		return
	}
	removeGuildConfig(g.ID)
	log.Infoln("Left guild: ", g.ID)
}

func removeGuildConfig(guildID string) {
	const query = `DELETE FROM core_config WHERE guild_id=$1`
	common.PQ.Exec(query, guildID)
}

func isGuildBanned(guildID string) bool {
	exists, err := models.BannedGuilds(qm.Where("guild_id = ?", guildID)).Exists(context.Background(), common.PQ)
	if err != nil {
		return false
	}

	return exists
}