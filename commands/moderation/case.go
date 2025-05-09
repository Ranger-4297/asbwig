package moderation

import (
	"context"
	"fmt"

	"github.com/RhykerWells/asbwig/bot/functions"
	"github.com/RhykerWells/asbwig/commands/moderation/models"
	"github.com/RhykerWells/asbwig/common"
	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type logAction struct {
	Name   string
	Colour int
}

var (
	logWarn   = logAction{Name: "Warned", Colour: 0xFCA253}
	logMute   = logAction{Name: "Muted", Colour: 0x5772BE}
	logUnmute = logAction{Name: "Unmuted", Colour: common.SuccessGreen}
	logKick   = logAction{Name: "Kicked", Colour: 0xF2A013}
	logBan    = logAction{Name: "Banned", Colour: 0xD64848}
	logUnban  = logAction{Name: "Unbanned", Colour: common.SuccessGreen}
)

const (
	caseEmoji    string = "<:ID:1369739780958457966>"
	userEmoji    string = "<:Member:1369740929568739499>"
	actionEmoji  string = "<:Action:1369745870001799321>"
	channelEmoji string = "<:Channel:1369743815887294687>"
	reasonEmoji  string = "<:Reason:1369744280310124624>"
)

// Case generation

// setupModerationCase returns the models.ModerationCase struct of data to insert into the database
func caseUpsert(caseID int64, guildID, staffID, targetID, reason, loglink string, action logAction) error {
	moderationCase := models.ModerationCase{
		CaseID: caseID,
		GuildID: guildID,
		StaffID: staffID,
		OffenderID: targetID,
		Reason: null.StringFrom(reason),
		Action: action.Name,
		Loglink: loglink,
	}
	
	return addCase(guildID, moderationCase)
}

// getNewCaseID returns the caseID for the currentCase
func getNewCaseID(guildID string) int64 {
	guild, _ := models.ModerationConfigs(qm.Where("guild_id=?", guildID)).One(context.Background(), common.PQ)
	currentCaseID := guild.LastCaseID + 1
	return currentCaseID
}

// incrementCaseID increments the guilds caseID. This is only used once everything else is successful. Such as posting the log and setting up the case data.
func incrementCaseID(guildID string) error {
	guild, _ := models.ModerationConfigs(qm.Where("guild_id=?", guildID)).One(context.Background(), common.PQ)
	guild.LastCaseID++
	_, err := guild.Update(context.Background(), common.PQ, boil.Infer())
	return err
}

// addCase adds the case to the database and runs incrementCaseID.
func addCase(guildID string, caseData models.ModerationCase) error {
	err := caseData.Insert(context.Background(), common.PQ, boil.Infer())
	if err != nil {
		return err
	}
	err = incrementCaseID(guildID)
	if err != nil {
		removeFailedCase(caseData)
		return err
	}
	return nil
}

func removeFailedCase(caseData models.ModerationCase) {
	caseData.Delete(context.Background(), common.PQ)
}

// Log generation

func logCase(guildID string, Author, Target *discordgo.Member, action logAction, currentChannel, reason string) error {
	logChannel, err := getGuildModLogChannel(guildID)
	if err != nil {
		return err
	}
	caseID := getNewCaseID(guildID)
	embed := logEmbed(Author.User, Target.User, caseID, action, currentChannel, reason)
	message, err := functions.SendMessage(logChannel, &discordgo.MessageSend{Embed: embed})
	if err != nil {
		return err
	}
	loglink := generateLogLink(guildID, logChannel, message.ID)
	caseUpsert(caseID, guildID, Author.User.ID, Target.User.ID, reason, loglink, action)
	return nil
}

// logEmbed returns the fully-populated embed for moderation logging
func logEmbed(author, target *discordgo.User, caseNumber int64, action logAction, channelID, reason string) *discordgo.MessageEmbed {
	humanReadableCaseNumber := humanize.Comma(caseNumber)
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    fmt.Sprintf("%s (ID %s)", author.Username, author.ID),
			IconURL: author.AvatarURL("1024"),
		},
		Description: fmt.Sprintf("%s **Case number:** %s\n%s **Who:** %s `(ID %s)`\n%s **Action:** %s\n%s **Channel:** <#%s>\n%s **Reason:** %s", caseEmoji, humanReadableCaseNumber, userEmoji, target.Mention(), target.ID, actionEmoji, action.Name, channelEmoji, channelID, reasonEmoji, reason),
		Color: action.Colour,
	}
}

// generateLogLink returns the full messageURL of the modlog entry
func generateLogLink(guildID, channelID, messageID string) string {
	link := fmt.Sprintf("https://discord.com/channels/%s/%s/%s", guildID, channelID, messageID)
	return link
}