package moderation

import (
	"context"
	"errors"

	"github.com/RhykerWells/asbwig/bot/functions"
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

var (
	errNotMuted = errors.New("user not muted")
	errNotMember = errors.New("user not a member")
)


func unmuteUser(guildID string, author, target string) error {
	moderationConfig, _ := models.ModerationConfigs(qm.Where("guild_id = ?", guildID)).One(context.Background(), common.PQ)
	muteUser, err := models.ModerationMutes(qm.Where("guild_id = ?", guildID), qm.Where("user_id = ?", target)).One(context.Background(), common.PQ)
	if err != nil {
		return errNotMuted
	}

	targetMember, err := functions.GetMember(guildID, target)
	if err != nil {
		if author == common.Bot.ID {
			muteUser.Delete(context.Background(), common.PQ)
		}
		return errNotMember
	}

	for _, roleID := range muteUser.Roles {
		functions.AddRole(guildID, target, roleID)
	}
	err = functions.RemoveRole(guildID, target, moderationConfig.MuteRole.String)
	if err != nil {
		return err
	}

	muteUser.Delete(context.Background(), common.PQ)

	if author == common.Bot.ID {
		modlogChannel, _ := getGuildModLogChannel(guildID)
		botMember, _ := functions.GetMember(guildID, common.Bot.ID)
		logCase(guildID, botMember, targetMember, logUnmute, modlogChannel, "Automatic unmute")
	}

	return nil
}