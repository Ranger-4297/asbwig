package deposit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/RhykerWells/asbwig/bot/functions"
	"github.com/RhykerWells/asbwig/commands/economy/models"
	"github.com/RhykerWells/asbwig/common"
	"github.com/RhykerWells/asbwig/common/dcommand"
	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var Command = &dcommand.AsbwigCommand{
	Command:     "deposit",
	Description: "Deposits a given amount into your bank",
	Args: []*dcommand.Args{
		{Name: "Amount", Type: dcommand.Int},
	},
	Run: func(data *dcommand.Data) {
		guild, _ := models.EconomyConfigs(qm.Where("guild_id=?", data.GuildID)).One(context.Background(), common.PQ)
		symbol := guild.Symbol
		userCash, err := models.EconomyCashes(qm.Where("guild_id=? AND user_id=?", data.GuildID, data.Author.ID)).One(context.Background(), common.PQ)
		var cash int64 = 0
		if err == nil {
			cash = userCash.Cash
		}
		userBank, err := models.EconomyBanks(qm.Where("guild_id=? AND user_id=?", data.GuildID, data.Author.ID)).One(context.Background(), common.PQ)
		var bank int64 = 0
		if err == nil {
			bank = userBank.Balance
		}
		embed := &discordgo.MessageEmbed{Author: &discordgo.MessageEmbedAuthor{Name: data.Author.Username, IconURL: data.Author.AvatarURL("256")}, Timestamp: time.Now().Format(time.RFC3339), Color: common.ErrorRed}
		if len(data.Args) <= 0 {
			embed.Description = "No `Amount` argument provided"
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		value := data.Args[0]
		if functions.ToInt64(value) <= 0 && strings.ToLower(value) != "all" {
			embed.Description = "Invalid `Amount` argument provided"
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		var amount int64
		if strings.ToLower(value) == "all" {
			amount = cash
		} else {
			amount = functions.ToInt64(value)
		}
		if amount > cash {
			embed.Description = fmt.Sprintf("You're unable to deposit more than you have in cash\nYou currently have %s%s", symbol, humanize.Comma(cash))
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		if cash < 0 {
			embed.Description = fmt.Sprintf("You're unable to deposit your overdraft\nYou are currently %s%s in arrears", symbol, humanize.Comma(cash))
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		userCash.Cash = cash - amount
		userBank.Balance = bank + amount
		_, _ = userCash.Update(context.Background(), common.PQ, boil.Whitelist("cash"))
		_, _ = userBank.Update(context.Background(), common.PQ, boil.Whitelist("balance"))
		embed.Description = fmt.Sprintf("You deposited %s%s into your bank\nThere is now %s%s in your bank", symbol, humanize.Comma(amount), symbol, humanize.Comma(userBank.Balance))
		embed.Color = common.SuccessGreen
		functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
	},
}
