package commands

import "github.com/bwmarrin/discordgo"

type CommandSignup struct{}

func (commandSignup *CommandSignup) Invokes() []string {
	return []string{"anmelden", "a"}
}

func (commandSignup *CommandSignup) AdminPermissionsNeeded() bool {
	return true
}

func (commandSignup *CommandSignup) Execute(ctx *Context) (err error) {
	embed := discordgo.MessageEmbed{
		Title:       "Beim Programmier-Wettbewerb anmelden",
		Description: "Du kannst dich ganz einfach anmelden! Gestatte mir, dich auf meine Weboberfläche weiterzuleiten.\n\n[***➤ HIER ANMELDEN***](https://google.de)",
		Color:       46074,
	}
	_, err = ctx.Session.ChannelMessageSendEmbed(ctx.Message.ChannelID, &embed)
	err = ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
	return
}
