package malo

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type command struct {
	app *cli.App
}

func New() *command {
	cmd := new(command)
	cmd.app = cli.NewApp()

	return cmd
}

func (c *command) Run() error {
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()
	return c.app.Run(os.Args)
}

func Run() {
	c := New()

	if err := c.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
