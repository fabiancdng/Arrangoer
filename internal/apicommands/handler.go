package apicommands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fabiancdng/Arrangoer/internal/config"
)

// Eine Funktion, die Befehle von der API entgegennimmt und diese ausführt
// Diese Funktion hat Zugriff auf den Botnutzer und kann diesen nutzen
// Dies sind z. B. Befehle wie eine Nachricht senden, etc.
func HandleAPICommand(session *discordgo.Session, config *config.Config, rawAPICommand string) {
	commandSlice := strings.Split(rawAPICommand, "///")

	invoke := strings.ToLower(commandSlice[0])
	args := commandSlice[1:]

	switch invoke {
	case "signup":
		message := fmt.Sprintf("**Eine Bewerbung von <@%s> ist soeben eingegangen!** 🥳\n\nDu wirst benachrichtigt, sobald deine Bewerbung angenommen oder abgelehnt wurde.", args[0])

		if args[1] != "" {
			message += fmt.Sprintf("\n\nSobald dein Team _'%s'_ vom Spielleiter bestätigt wurde, wirst du benachrichtigt und bekommst die Discord-Rolle automatisch zugewiesen.", args[1])
		}

		embed := &discordgo.MessageEmbed{
			Title:       "Bewerbung eingegangen",
			Description: message,
			Color:       58176,
		}
		session.ChannelMessageSendEmbed(config.Discord.NotificationsChannelID, embed)
	}
}
