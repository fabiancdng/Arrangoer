package apicommands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Eine Funktion, die Befehle von der API entgegennimmt und diese ausführt
// Diese Funktion hat Zugriff auf den Botnutzer und kann diesen nutzen
// Dies sind z. B. Befehle wie eine Nachricht senden, etc.
func HandleAPICommand(session *discordgo.Session, rawAPICommand string) {
	commandSlice := strings.Split(rawAPICommand, "///")

	invoke := strings.ToLower(commandSlice[0])
	args := commandSlice[1:]

	log.Println(invoke, args)
}
