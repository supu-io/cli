package main

import (
	"encoding/json"
	"log"
	"os"
	"os/user"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

// Config is the configuration struct for this app
type Config struct {
	Supu struct {
		URL string `json:"url"`
	} `json:"supu-io"`
}

// Get config defined on the .supu file
func getConfig() *Config {
	source := getConfigPath()
	c := Config{}
	file, err := os.Open(source)
	if err != nil {
		log.Panic("error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Println("Config file is invalid")
		log.Panic("error:", err)
	}
	return &c
}

// Get the config path to use, default is .supu on the same
// folder, but you can override it with a same named file on
// your home
func getConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir + "/.supu"

	if _, err := os.Stat(home); err == nil {
		color.Green("Home found")
		return home
	}
	return ".supu"
}

func main() {
	config := getConfig()
	m := Manager{URL: config.Supu.URL}
	app := cli.NewApp()
	app.Name = "supu"
	app.Usage = "make an explosive entrance"
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list :status:",
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					color.Red("You should specify the status")
				} else {
					status := c.Args()[0]
					m.list(status)
				}
			},
		},
		{
			Name:    "show",
			Aliases: []string{"l"},
			Usage:   "show :issue:",
			Action: func(c *cli.Context) {
				issue := c.Args()[1]
				m.details(issue)
			},
		},
		{
			Name:    "comment",
			Aliases: []string{"l"},
			Usage:   "comment :issue: :body:",
			Action: func(c *cli.Context) {
				issue := c.Args()[1]
				body := c.Args()[2]
				m.comment(issue, body)
			},
		},
		{
			Name:    "move",
			Aliases: []string{"m"},
			Usage:   "move [back|forward] issue",
			Subcommands: []cli.Command{
				{
					Name:  "back",
					Usage: "Move the issue back on the workflow",
					Action: func(c *cli.Context) {
						issue := c.Args()[0]
						m.move(issue, "back")
					},
				},
				{
					Name:  "forward",
					Usage: "Move the issue back on the workflow",
					Action: func(c *cli.Context) {
						issue := c.Args()[0]
						m.move(issue, "forward")
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
