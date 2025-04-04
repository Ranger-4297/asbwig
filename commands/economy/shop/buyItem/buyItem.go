package buyitem

import (
	"context"
	"fmt"
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
	Command:     "buyitem",
	Description: "Buys an item from the shop",
	Args: []*dcommand.Args{
		{Name: "Name", Type: dcommand.String},
		{Name: "Quantity", Type: dcommand.Int},
	},
	Run: func(data *dcommand.Data) {
		guild, _ := common.Session.Guild(data.GuildID)
		embed := &discordgo.MessageEmbed{Author: &discordgo.MessageEmbedAuthor{Name: guild.Name + " Store", IconURL: guild.IconURL("256")}, Timestamp: time.Now().Format(time.RFC3339), Color: common.ErrorRed}
		guildConfig, _ := models.EconomyConfigs(qm.Where("guild_id=?", data.GuildID)).One(context.Background(), common.PQ)
		economyUser, err := models.EconomyUsers(qm.Where("guild_id=? AND user_id=?", data.GuildID, data.Author.ID)).One(context.Background(), common.PQ)
		var cash int64 = 0
		if err == nil {
			cash = economyUser.Cash
		}
		economyUserInventory, _ := models.EconomyUserInventories(qm.Where("guild_id=? AND user_id=?", data.GuildID, data.Author.ID)).All(context.Background(), common.PQ)
		if len(data.Args) <= 0 {
			embed.Description = "No `Item` argument provided\nUse `shop [Page]` to view all items"
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		name := data.ArgsNotLowered[0]
		item, exists := models.EconomyShops(qm.Where("guild_id=? AND name=?", data.GuildID, name)).One(context.Background(), common.PQ)
		if exists != nil {
			embed.Description = "This item doesn't exist"
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		var buyQuantity int64 = 1
		if len(data.Args) > 1 {
			if data.Args[1] == "max" || data.Args[1] == "all" {
				quantity := map[string]int64{"max": (cash / item.Price), "all": item.Quantity}
				buyQuantity=quantity[data.Args[1]]
				if buyQuantity == 0 {
					buyQuantity=quantity["max"]
				}
			} else if functions.ToInt64(data.Args[1]) > 0 {
				buyQuantity = functions.ToInt64(data.Args[1])
			}
		}
		if item.Quantity > 0 && buyQuantity > item.Quantity {
			embed.Description = "There's not enough of this in the shop to buy that much"
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		if (name == "Chicken" || name == "chicken") && buyQuantity > 1 {
			embed.Description = "You can't buy more than one chicken at a time"
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		var chicken bool = false
		for _, inventoryItem := range economyUserInventory {
			if inventoryItem.Name == "Chicken" || inventoryItem.Name == "chicken" {
				chicken = true
				break
			}
		}
		if chicken {
			embed.Description = "You can't have more than one chicken in your inventory"
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		if item.Soldby.String ==  data.Author.ID {
			embed.Description = "You can't buy an item that you have listed"
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		if (item.Price * buyQuantity) > cash {
			embed.Description = "You don't have enough money to buy this item"
			functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
			return
		}
		newQuantity := item.Quantity - buyQuantity
		if item.Quantity > 0 && newQuantity == 0 {
			item.Delete(context.Background(), common.PQ)
		} else if item.Quantity != 0 {
			item.Quantity = newQuantity
			item.Update(context.Background(), common.PQ, boil.Infer())
		}
		newQuantity = buyQuantity
		for _, inventoryItem := range economyUserInventory {
			if inventoryItem.Name == name {
				newQuantity = inventoryItem.Quantity + buyQuantity
				break
			}
		}
		embed.Color = common.SuccessGreen
		embed.Description = fmt.Sprintf("You bought %s of %s for %s%s", humanize.Comma(buyQuantity), name, guildConfig.Symbol, humanize.Comma(item.Price * buyQuantity))
		userInventory := models.EconomyUserInventory{GuildID: data.GuildID, UserID: data.Author.ID, Name: name, Quantity: newQuantity, Role: item.Role, Reply: item.Reply}
		_ = userInventory.Upsert(context.Background(), common.PQ, true, []string{"guild_id", "user_id", "name"}, boil.Whitelist("quantity"), boil.Infer())
		cash = cash - (item.Price * buyQuantity)
		userEntry := models.EconomyUser{GuildID: data.GuildID, UserID: data.Author.ID, Cash: cash}
		_ = userEntry.Upsert(context.Background(), common.PQ, true, []string{"guild_id", "user_id"}, boil.Whitelist("cash"), boil.Infer())
		functions.SendMessage(data.ChannelID, &discordgo.MessageSend{Embed: embed})
	},
}