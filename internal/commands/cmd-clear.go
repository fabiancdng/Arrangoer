package commands

type CommandClear struct{}

func (commandClear *CommandClear) Invokes() []string {
	return []string{"test"}
}

func (commandClear *CommandClear) AdminPermissionsNeeded() bool {
	return true
}

func (CommandClear *CommandClear) Execute(ctx *Context) (err error) {
	err = ctx.Session.MessageReactionAdd(ctx.Message.ChannelID, ctx.Message.ID, "✅")
	return
}
