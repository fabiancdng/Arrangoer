package commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	prefix      string
	cmdMap      map[string]Command
	middlewares []Middleware
}

func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix:      prefix,
		cmdMap:      make(map[string]Command),
		middlewares: make([]Middleware, 0),
	}
}

func (commandHandler *CommandHandler) RegisterCommand(command Command) {
	for _, invoke := range command.Invokes() {
		commandHandler.cmdMap[invoke] = command
	}
}

func (commandHandler *CommandHandler) RegisterMiddleware(mw Middleware) {
	commandHandler.middlewares = append(commandHandler.middlewares, mw)
}

func (commandHandler *CommandHandler) HandleMessage(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.Bot {
		return
	}

	if !strings.HasPrefix(event.Content, commandHandler.prefix) {
		return
	}

	// Nachricht: <Prefix><Invoke> <Arg 1> <Arg2> ...
	// messageSlice: [<invoke>, <Arg1>, <Arg2>, ...]
	messageSlice := strings.Split(event.Content[len(commandHandler.prefix):], " ")
	invoke := strings.ToLower(messageSlice[0])
	args := messageSlice[1:]

	// 'valid' ist ein Boolean, der anzeigt, ob der Command in der Map existiert
	command, valid := commandHandler.cmdMap[invoke]
	if !valid || command == nil {
		return
	}

	ctx := &Context{
		Session: session,
		Args:    args,
		Handler: commandHandler,
		Message: event.Message,
	}

	// Alle Middlewares vor der Command-Ausführung ausführen
	for _, mw := range commandHandler.middlewares {
		next, err := mw.Execute(ctx, command)
		if err != nil {
			return
		}

		if !next {
			return
		}
	}

	// Den Command ausführen
	err := command.Execute(ctx)
	if err != nil {
		return
	}

	log.Printf("%s hat den Befehl '%s' ausgeführt.", ctx.Message.Author.Username, invoke)
	return

}
