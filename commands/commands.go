package commands

import (
	"github.com/RhykerWells/asbwig/common/dcommand"
	"github.com/bwmarrin/discordgo"

	"github.com/RhykerWells/asbwig/commands/economy"
	"github.com/RhykerWells/asbwig/commands/moderation"
	
	"github.com/RhykerWells/asbwig/commands/standard/invite"
	"github.com/RhykerWells/asbwig/commands/standard/ping"

	"github.com/RhykerWells/asbwig/commands/botOwner/setstatus"
)

var cmdHandler *dcommand.CommandHandler

func InitCommands(session *discordgo.Session) {
	cmdHandler = dcommand.NewCommandHandler()

	cmdHandler.RegisterCommands(
		helpCmd,
		prefixCmd,

		ping.Command,
		invite.Command,

		setstatus.Command,
	)

	economy.EconomySetup(cmdHandler)
	moderation.ModerationSetup(cmdHandler)
	session.AddHandler(cmdHandler.HandleMessageCreate)
}