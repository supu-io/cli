package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

type Config struct {
	Supu struct {
		Url string `json:"url"`
	} `json:"supu-io"`
}

func getConfig(source string) *Config {
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

func main() {
	config := getConfig("config.json")
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) {
		m := Manager{Context: c, Url: config.Supu.Url}
		m.Manage()
	}

	app.Run(os.Args)
}
