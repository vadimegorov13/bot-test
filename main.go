package main

import (
	"GoBot/bot"
	"GoBot/config"
	"fmt"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println("ERROR reading config.json", err)
		return
	}

	bot.Start()

}
