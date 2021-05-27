package apicommands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Simpler switch-basierter Command-Handler, der Invokes von der API-
// Goroutine entgegennimmt und durch Zugriff auf den Bot-Nutzer
// z. B. Dinge machen kann wie Nachrichten senden oder Rollen erstellen
func HandleAPICommand(ctx *Context) {
	commandSlice := strings.Split(ctx.Command, "///")

	invoke := strings.ToLower(commandSlice[0])
	args := commandSlice[1:]

	switch invoke {
	// Eine Anmeldung ist eingegangen
	case "signup":
		message := fmt.Sprintf("**Eine Anmeldung von <@%s> ist soeben eingegangen!** 🥳\n\nDu wirst benachrichtigt, sobald deine Anmeldung angenommen oder abgelehnt wurde.", args[0])

		if args[1] != "" {
			message += fmt.Sprintf("\n\nSobald dein Team **%s** vom Spielleiter bestätigt wurde, wirst du benachrichtigt und bekommst die Discord-Rolle automatisch zugewiesen.", args[1])
		}

		embed := &discordgo.MessageEmbed{
			Title:       "Anmeldung eingegangen",
			Description: message,
			Color:       58176,
		}
		ctx.Session.ChannelMessageSendEmbed(ctx.Config.Discord.NotificationsChannelID, embed)

	// Eine Anmeldung wurde akzeptiert
	case "signup-accepted":
		message := fmt.Sprintf("**Die Anmeldung von <@%s> wurde soeben akzeptiert!** 🥳", args[0])

		embed := &discordgo.MessageEmbed{
			Title:       "Anmeldung akzeptiert",
			Description: message,
			Color:       58176,
		}
		ctx.Session.ChannelMessageSendEmbed(ctx.Config.Discord.NotificationsChannelID, embed)
	}
}
