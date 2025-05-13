package moderation

import (
	"context"
	"strings"
	"time"

	"slices"

	"github.com/RhykerWells/asbwig/bot/functions"
	"github.com/RhykerWells/asbwig/commands/moderation/models"
	"github.com/RhykerWells/asbwig/common"
	"github.com/RhykerWells/asbwig/common/dcommand"
	"github.com/bwmarrin/discordgo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// returns a boolean value on the status of the guild moderation
func isEnabled(guildID string) bool {
	config, _ := models.ModerationConfigs(qm.Where("guild_id=?", guildID)).One(context.Background(), common.PQ)
	return config.Enabled
}

// returns an array of required roles to run the selected command
func requireRoles(guildID, command string) types.StringArray {
	config, _ := models.ModerationConfigs(qm.Where("guild_id=?", guildID)).One(context.Background(), common.PQ)
	var requiredRoles types.StringArray
	switch command{
	case "warn":
		requiredRoles = config.RequiredWarnRoles
	case "mute":
		requiredRoles = config.RequiredMuteRoles
	case "unmute":
		requiredRoles = config.RequiredUnmuteRoles
	case "kick":
		requiredRoles = config.RequiredKickRoles
	case "ban":
		requiredRoles = config.RequiredBanRoles
	case "unban":
		requiredRoles = config.RequiredUnbanRoles
	}
	return requiredRoles
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
		if !enabled {
			functions.DeleteMessage(data.ChannelID, data.Message.ID, 1*time.Second)
			functions.SendBasicMessage(data.ChannelID, "The moderation system has not been enabled please enable it on the dashboard.", 30*time.Second)
			return
		}
		author, _ := functions.GetMember(data.GuildID, data.Author.ID)
		ok := hasCommandPermissions(data.GuildID, author, "warn")
		if !ok {
			functions.DeleteMessage(data.ChannelID, data.Message.ID, 1*time.Second)
			functions.SendBasicMessage(data.ChannelID, "You don't have the required roles for this command.", 30*time.Second)
			return
		}
		target, _ := functions.GetMember(data.GuildID, data.Args[0])
		if target.User.ID == author.User.ID {
			functions.DeleteMessage(data.ChannelID, data.Message.ID, 1*time.Second)
			functions.SendBasicMessage(data.ChannelID, "You can't warn yourself.", 30*time.Second)
			return
		}
		ok = functions.IsMemberHigher(data.GuildID, author, target)
		if !ok {
			functions.DeleteMessage(data.ChannelID, data.Message.ID, 1*time.Second)
			functions.SendBasicMessage(data.ChannelID, "You don't have the correct permissions to warn this user (target has higher role).")
			return
		}
		warnReason := strings.Join(data.Args[1:], " ")
		err := logCase(data.GuildID, author, target, logWarn, data.ChannelID, warnReason)
		if err != nil {
			functions.SendBasicMessage(data.ChannelID, "Please setup a modlog channel before running this command", 30*time.Second)
			return
		}
	},
}