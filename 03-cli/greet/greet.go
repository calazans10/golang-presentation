package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/calazans10/hello"
	"github.com/urfave/cli"
)

func main() {
	var language string
	var uppercase bool

	app := cli.NewApp()
	app.Name = "greet"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Jeferson Farias Calazans",
			Email: "calazans10@gmail.com",
		},
	}
	app.Usage = "fight the loneliness!"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "upper, u",
			Usage:       "put the greeting in capital letters",
			Destination: &uppercase,
		},
		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
	}

	app.Action = func(c *cli.Context) error {
		name := ""
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		greeting := hello.Greeting(name, strings.Title(language))
		if uppercase {
			greeting = strings.ToUpper(greeting)
		}
		fmt.Println(greeting)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
