package dcommand

import "github.com/bwmarrin/discordgo"

type Data struct {
	Session *discordgo.Session

	GuildID   string // GuildID of the executed command
	ChannelID string // ChannelID the command was executed in
	Author    *discordgo.User // User object of the triggering user

	Message        *discordgo.Message // Message object of the triggering message
	Args           []string // Lowered args parsed to commands
	ArgsNotLowered []string // Args parsed to command

	Handler *CommandHandler
}