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
	logKick   = logAction{Name: "Kicked", Colour: 0xFCA253}
	logBan    = logAction{Name: "Banned", Colour: common.SuccessGreen}
	loguUban  = logAction{Name: "Unbanned", Colour: 0xFCA253}
)

const (
	caseEmoji    string = "<:ID:1369739780958457966>"
	userEmoji    string = "<:Member:1369740929568739499>"
	actionEmoji  string = "<:Action:1369745870001799321>"
	channelEmoji string = "<:Channel:1369743815887294687>"
	reasonEmoji  string = "<:Reason:1369744280310124624>"
)

// logEmbed returns the fully-populated embed for moderation logging
func logEmbed(author, user *discordgo.User, caseNumber int64, action logAction, channelID, reason string) *discordgo.MessageEmbed {
	humanReadableCaseNumber := humanize.Comma(caseNumber)
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    fmt.Sprintf("%s (ID %s)", author.Username, author.ID),
			IconURL: author.AvatarURL("1024"),
		},
		Description: fmt.Sprintf("%s **Case number:** %s\n%s **Who:** %s `(ID %s)`\n%s **Action:** %s\n%s **Channel:** <#%s>\n%s **Reason:** %s", caseEmoji, humanReadableCaseNumber, userEmoji, user.Mention(), user.ID, actionEmoji, action.Name, channelEmoji, channelID, reasonEmoji, reason),
		Color: action.Colour,
	}
}

func logCase(guildID string, Author, Offender discordgo.Member, reason string, action logAction, currentChannel string) error {
	var logChannel string
	query := `SELECT mod_log FROM moderation_config WHERE guild_id=$1`

	err := common.PQ.QueryRow(query, guildID).Scan(&logChannel)
	if err != nil {
		return err
	}

	caseData := setupModerationCase(guildID, Author.User.ID, Offender.User.ID, reason, action)
	embed := logEmbed(Author.User, Offender.User, caseData.CaseID, action, currentChannel, reason)
	message, err := functions.SendMessage(logChannel, &discordgo.MessageSend{Embed: embed})
	if err != nil {
		removeErroneousCase(guildID, caseData.CaseID)
		return err
	}
	logLinkToCase(caseData, message.ChannelID, message.ID)
	return nil
}

func setupModerationCase(guildID, staffID, offenderID, reason string, action logAction) (caseData models.ModerationCase) {
	currentCaseNumber := getCurrentCaseNumber(guildID) + 1
	moderationCase := models.ModerationCase{
		CaseID: currentCaseNumber,
		GuildID: guildID,
		StaffID: staffID,
		OffenderID: offenderID,
		Reason: null.StringFrom(reason),
		Action: action.Name,
	}

	moderationCase.Insert(context.Background(), common.PQ, boil.Infer())
	return moderationCase
}

func getCurrentCaseNumber(guildID string) int64 {
	var currentCaseID int64
	query := `SELECT COALESCE(MAX(case_id), 0) FROM moderation_cases WHERE guild_id=$1`
	err := common.PQ.QueryRow(query, guildID).Scan(&currentCaseID)
	if err != nil {
		// No cases
		currentCaseID = 0
	}
	return currentCaseID
}

func logLinkToCase(caseData models.ModerationCase, channelID, messageID string) {
	link := generateLogLink(caseData.GuildID, channelID, messageID)
	caseData.Loglink = null.StringFrom(link)
	caseData.Update(context.Background(), common.PQ, boil.Infer())
}

func generateLogLink(guildID, channelID, messageID string) string {
	link := fmt.Sprintf("https://discord.com/channels/%s/%s/%s", guildID, channelID, messageID)
	return link
}

func removeErroneousCase(guildID string, caseID int64) {

}