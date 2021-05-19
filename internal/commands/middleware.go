package commands

type Middleware interface {
	Execute(ctx *Context, cmd Command) (next bool, err error)
}
