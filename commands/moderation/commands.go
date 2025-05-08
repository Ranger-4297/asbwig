package moderation

import (
	"strings"
	"time"

	"github.com/RhykerWells/asbwig/bot/functions"
	"github.com/RhykerWells/asbwig/common/dcommand"
)

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
	Run: (func(data *dcommand.Data) {
		author, _ := functions.GetMember(data.GuildID, data.Author.ID)
		offender, _ := functions.GetMember(data.GuildID, data.Args[0])
		warnReason := strings.Join(data.Args[1:], " ")
		err := logCase(data.GuildID, *author, *offender, logWarn, data.ChannelID, warnReason)
		if err != nil {
			functions.SendBasicMessage(data.ChannelID, "Please setup a modlog channel before running this command", 30*time.Second)
			return
		}
	}),
}