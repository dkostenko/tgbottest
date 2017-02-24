package main

import (
	"flag"

	"github.com/dkostenko/tgbottest/app"
	"github.com/dkostenko/tgbottest/tg"
	ivona "github.com/jpadilla/ivona-go"
)

func main() {
	tgToken := flag.String("tg-token", "", "Telegram bot token")
	ivonaAK := flag.String("ivona-ak", "", "Ivona access key")
	ivonaSK := flag.String("ivona-sk", "", "Ivona secret key")
	flag.Parse()

	ivonaClient := ivona.New(*ivonaAK, *ivonaSK)
	tgClient := tg.NewClient(*tgToken)

	appManager := app.NewManager(tgClient, ivonaClient)
	appManager.Listen()
}
