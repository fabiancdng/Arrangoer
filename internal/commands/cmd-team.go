package commands

import "github.com/bwmarrin/discordgo"

type CommandTeam struct{}

func (commandTeam *CommandTeam) Invokes() []string {
	return []string{"team", "t"}
}

func (commandTeam *CommandTeam) AdminPermissionsNeeded() bool {
	return false
}

func (commandTeam *CommandTeam) Execute(ctx *Context) (err error) {
	embed := discordgo.MessageEmbed{
		Title:       "Zuweisung deines Teams",
		Description: "Bist du schon angemeldet und hast ein Team gefunden? Dann klicke einmal hier, damit du die entsprechende Rolle auf dem Discord bekommst.\n\n[***➤ HIER DEIN TEAM ZUWEISEN***](https://google.de)",
		Color:       14680128,
	}
	_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, &embed)
	return
}
