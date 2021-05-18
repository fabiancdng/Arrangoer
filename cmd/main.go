/*
								Arrangør

	Discord Bot, der bei der Organisation des Programmier-Wettbewerbs hilft.

	Einsendung für den 'Programmier-Wettbewerb' der 'Digitalen Woche 2021 Leer'

					Copyright (c) 2021 Fabian Reinders

*/

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/fabiancdng/Arrangoer/internal/commands"
	"github.com/fabiancdng/Arrangoer/internal/config"
	"github.com/fabiancdng/Arrangoer/internal/events"
)

func main() {
	config, err := config.ParseConfig("./config/config.json")
	if err != nil {
		log.Panic(err)
	}

	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Panic(err)
	}

	session.Identify.Intents = discordgo.IntentsAll

	registerEvents(session)
	registerCommands(session, config)

	err = session.Open()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Der Bot läuft jetzt! // Er kann mit STRG+C beendet werden.")

	sessionChannel := make(chan os.Signal, 1)
	signal.Notify(sessionChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sessionChannel

	session.Close()
}

func registerEvents(session *discordgo.Session) {
	session.AddHandler(events.NewReadyHandler().Handler)
	session.AddHandler(events.NewJoinHanlder().Handler)
}

func registerCommands(session *discordgo.Session, config *config.Config) {
	commandHandler := commands.NewCommandHandler(config.Prefix)

	commandHandler.RegisterCommand(&commands.CommandTest{})
	commandHandler.RegisterMiddleware(&commands.MiddlewarePermissions{})

	session.AddHandler(commandHandler.HandleMessage)
}
