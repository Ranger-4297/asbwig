package commands

import (
	"fmt"
	"strings"

	"github.com/RhykerWells/asbwig/bot/functions"
	"github.com/RhykerWells/asbwig/common"
	"github.com/RhykerWells/asbwig/common/dcommand"
	"github.com/bwmarrin/discordgo"
)

var helpCmd = &dcommand.AsbwigCommand{
	Command: "help",
	Aliases: []string{"h"},
	Args: []*dcommand.Args{
		{Name: "Command", Type: dcommand.String, Optional: true},
	},
	Description: "Displays bot help",
	Run:         helpFunc,
}

func helpFunc(data *dcommand.Data) {
	command := ""
	if len(data.Args) > 0 {
		command = data.Args[0]
	}

	// Per-command help
	if command != "" {
		help(command, data.ChannelID)
		return
	}

	// Generic help category
	basicEmbed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    fmt.Sprintf("%s help", common.Bot.Username),
			IconURL: common.Bot.AvatarURL("256"),
		},
		Description: "Hi there, I'm ASBWIG. As of right now I am vastly unpopulated.\nShould probably add somewhere for you to view all commands huh..",
		Color:       common.SuccessGreen,
	}
	message := &discordgo.MessageSend{
		Embed: basicEmbed,
	}
	functions.SendMessage(data.ChannelID, message)
}

func help(command string, channelID string) {
	cmdMap := dcommand.CmdHndlr.RegisteredCommands()
	cmd, ok := cmdMap[command]
	if !ok {
		functions.SendBasicMessage(channelID, fmt.Sprintf("Command `%s` not found", command))
		return
	}
	helpEmbed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    fmt.Sprintf("%s help - %s/%s", common.Bot.Username, command, strings.Join(cmd.Aliases, "/")),
			IconURL: common.Bot.AvatarURL("256"),
		},
		Description: cmd.Description,
		Color:       common.SuccessGreen,
	}
	args := getArgs(cmd)
	helpEmbed.Description = cmd.Description
	if args != "" {
		helpEmbed.Description += "\n```" + cmd.Trigger + args + "\n```"
	}
	message := &discordgo.MessageSend{
		Embed: helpEmbed,
	}
	functions.SendMessage(channelID, message)
}

func getArgs(command dcommand.RegisteredCommand) (str string) {
	for _, arg := range command.Args {
		if arg.Optional {
			str += " [" + argHelp(arg) + "]"
		} else {
			str += " <" + argHelp(arg) + ">"
		}
	}
	return
}

func argHelp(arg *dcommand.Args) (str string) {
	argType := arg.Type.Help()
	str = fmt.Sprintf("%s:%s", arg.Name, argType)
	return
}
