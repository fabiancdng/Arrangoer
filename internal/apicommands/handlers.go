package apicommands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fabiancdng/Arrangoer/internal/models"
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
		message := fmt.Sprintf("**Eine Anmeldung von <@!%s> ist soeben eingegangen!** 🥳\n\nDu wirst benachrichtigt, sobald deine Anmeldung angenommen oder abgelehnt wurde.", args[0])

		if args[1] != "" {
			message += fmt.Sprintf("\n\nSobald dein Team **%s** vom Spielleiter bestätigt wurde, wirst du benachrichtigt und bekommst die Discord-Rolle automatisch zugewiesen.", args[1])
		}

		embed := &discordgo.MessageEmbed{
			Title:       "Anmeldung eingegangen",
			Description: message,
			Color:       15204542,
		}
		ctx.Session.ChannelMessageSendEmbed(ctx.Config.Discord.NotificationsChannelID, embed)

	// Eine Anmeldung wurde akzeptiert
	case "signup-accepted":
		applicationID, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}

		application := new(models.Application)
		application, err = ctx.Db.GetApplication(applicationID)
		if err != nil {
			return
		}

		message := fmt.Sprintf("**Die Anmeldung von <@!%s> wurde soeben akzeptiert!** 🥳\n\nFalls dein Team noch nicht bestätigt wurde, folgt eine Benachrichtigung sowie eine automatische Zuweisung der Rolle noch 😊", application.UserID)

		embed := &discordgo.MessageEmbed{
			Title:       "Anmeldung akzeptiert",
			Description: message,
			Color:       62781,
		}

		ctx.Session.ChannelMessageSendEmbed(ctx.Config.Discord.NotificationsChannelID, embed)
	}
}
