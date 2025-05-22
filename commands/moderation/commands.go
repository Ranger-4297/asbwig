package moderation

import (
	"context"
	"fmt"
	"strings"
	"time"

	"slices"

	"github.com/RhykerWells/asbwig/bot/functions"
	"github.com/RhykerWells/asbwig/commands/moderation/models"
	"github.com/RhykerWells/asbwig/common"
	"github.com/RhykerWells/asbwig/common/dcommand"
	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// returns a boolean value on the status of the guild moderation
func isEnabled(guildID string) bool {
	config, _ := models.ModerationConfigs(qm.Where("guild_id=?", guildID)).One(context.Background(), common.PQ)
	return config.Enabled
}

// returns an array of required roles to run the selected command
func requireRoles(guildID, command string) []string {
	requiredRoles, _ := models.ModerationConfigRoles(qm.Where("guild_id = ?", guildID), qm.Where("action_type = ?", command)).All(context.Background(), common.PQ)
	var roleIDs []string
	for _, role := range requiredRoles {
		roleIDs = append(roleIDs, role.RoleID)
	}

	return roleIDs
}

// returns a boolean on whether the user has the current permissions to run the selected command
func hasCommandPermissions(guildID string, user *discordgo.Member, requireType string) bool {
	requiredRoles := requireRoles(guildID, requireType)
	for _, role := range user.Roles {
		if slices.Contains(requiredRoles, role) {
			return true
		}
	}
	return false
}

func triggerDeltion(guildID string) (bool, int) {
	config, _ := models.ModerationConfigs(qm.Where("guild_id=?", guildID)).One(context.Background(), common.PQ)
	return config.EnabledTriggerDeletion, config.SecondsToDeleteTrigger
}

func responseDeletion(guildID string) (bool, int) {
	config, _ := models.ModerationConfigs(qm.Where("guild_id=?", guildID)).One(context.Background(), common.PQ)
	return config.EnabledResponseDeletion, config.SecondsToDeleteResponse
}

// response returns the fully-populated embed for responses
func responseEmbed(author, target *discordgo.User, action logAction) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    fmt.Sprintf("Case type: %s", action.CaseType),
			IconURL: author.AvatarURL("1024"),
		},
		Description: fmt.Sprintf("%s has successfully %s %s :thumbsup:", author.Mention(), action.Name, target.Mention()),
		Color: 0x242429,
	}
}

var warnCommand = &dcommand.AsbwigCommand{
	Command:     "warn",
	Category: 	 dcommand.CategoryModeration,
	Aliases:     []string{""},
	Description: "Warns a user for a specified reason",
	Args: []*dcommand.Args{
		{Name: "User", Type: dcommand.User},
		{Name: "Reason", Type: dcommand.String},
	},
	ArgsRequired: 2,
	Run: func(data *dcommand.Data) {
		enabled := isEnabled(data.GuildID)
		guild := functions.GetGuild(data.GuildID)
		if !enabled {
			functions.DeleteMessage(data.ChannelID, data.Message.ID, 1*time.Second)
			functions.SendBasicMessage(data.ChannelID, "The moderation system has not been enabled please enable it on the dashboard.", 10*time.Second)
			return
		}
		author, _ := functions.GetMember(data.GuildID, data.Author.ID)
		ok := hasCommandPermissions(data.GuildID, author, "Warn")
		if !ok {
			functions.DeleteMessage(data.ChannelID, data.Message.ID, 1*time.Second)
			functions.SendBasicMessage(data.ChannelID, "You don't have the required roles for this command.", 10*time.Second)
			return
		}
		target, _ := functions.GetMember(data.GuildID, data.Args[0])
		if target.User.ID == author.User.ID {
			functions.DeleteMessage(data.ChannelID, data.Message.ID, 1*time.Second)
			functions.SendBasicMessage(data.ChannelID, "You can't warn yourself.", 10*time.Second)
			return
		}
		ok = functions.IsMemberHigher(data.GuildID, author, target)
		if (!ok || target.User.ID == guild.OwnerID) {
			functions.DeleteMessage(data.ChannelID, data.Message.ID, 1*time.Second)
			functions.SendBasicMessage(data.ChannelID, "You don't have the correct permissions to warn this user (target has higher role).", 10*time.Second)
			return
		}
		warnReason := strings.Join(data.Args[1:], " ")
		err := logCase(data.GuildID, author, target, logWarn, data.ChannelID, warnReason)
		if err != nil {
			functions.SendBasicMessage(data.ChannelID, "Please setup a modlog channel before running this command", 10*time.Second)
			return
		}
		ok, delay := triggerDeltion(data.GuildID)
		if ok {
			functions.DeleteMessage(data.ChannelID, data.Message.ID, time.Duration(delay)*time.Second)
		}
		responseEmbed := responseEmbed(author.User, target.User, logWarn)
		message, _ := functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: responseEmbed})
		ok, delay = responseDeletion(data.GuildID)
			if ok {
			functions.DeleteMessage(data.ChannelID, message.ID, time.Duration(delay)*time.Second)
		}
	},
}