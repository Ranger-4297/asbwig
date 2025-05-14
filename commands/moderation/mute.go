package moderation

import (
	"context"

	"github.com/RhykerWells/asbwig/commands/moderation/models"
	"github.com/RhykerWells/asbwig/common"
	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func RefreshMuteSettings(guildID string) {
	config, err := models.ModerationConfigs(qm.Where("guild_id = ?", guildID)).One(context.Background(), common.PQ)
	if err != nil {
		return
	}
	managed := config.ManageMuteRole
	if !managed {
		return
	}
	if config.MuteRole.String == "" {
		return
	}
	channels, _ := common.Session.GuildChannels(guildID)
	for _, channel := range channels {
		common.Session.ChannelPermissionSet(channel.ID, config.MuteRole.String,  discordgo.PermissionOverwriteTypeRole, 0, discordgo.PermissionSendMessages)
	}
}