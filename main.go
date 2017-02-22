package main

import (
	"flag"
	"log"

	"github.com/dkostenko/tgbottest/tg"
)

func main() {
	token := flag.String("token", "", "Telegram bot token")
	flag.Parse()

	bot := tg.NewClient(*token)
	msgChan, _ := bot.GetMessagesChan()
	for msg := range msgChan {
		log.Println(msg)
	}
}
