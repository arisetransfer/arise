package main

import (
	"os"
	"log"

	"github.com/urfave/cli/v2"
	"github.com/arisetransfer/arise/client"
	"github.com/arisetransfer/arise/server"
)

func main() {
	app := &cli.App{
		Name: "arise",
		Usage: "Transfer file between two devices",
	}
	app.Commands = []*cli.Command{
		{
			Name: "send",
			ArgsUsage: "[filename]",
			Usage: "Send the file over relay",
			HelpName: "arise send",
			Action: func (c *cli.Context) error {
				client.Sender(c.Args().Get(0))
				return nil
			},
		},
		{
			Name: "relay",
			ArgsUsage: "[Port]",
			Usage: "Start an arise relay on port default(6969)",
			Action: func (c *cli.Context) error {
				if c.Args().Get(0)=="" {
					server.StartRelay("6969")
					return nil
				}
				server.StartRelay(c.Args().Get(0))
				return nil
			},
		},
		{
			Name: "receive",
			ArgsUsage: "[unique_code]",
			Usage: "Receive file using the unique code",
			Action: func (c *cli.Context) error {
				if c.Args().Get(0)=="" {
					cli.ShowAppHelp(c)
					return nil
				}
				client.Reciever(c.Args().Get(0))
				return nil
			},
		},
	}
	err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
