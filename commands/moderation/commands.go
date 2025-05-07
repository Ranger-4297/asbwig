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
		err := logCase(data.GuildID, *author, *offender, warnReason, logWarn, data.ChannelID)
		if err != nil {
			functions.SendBasicMessage(data.ChannelID, "I can't access or send messages in the modlog channel. Can't run command", 30*time.Second)
			return
		}
	}),
}