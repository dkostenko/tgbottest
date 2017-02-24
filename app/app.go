package app

import (
	"log"

	"github.com/dkostenko/tgbottest/tg"
	ivona "github.com/jpadilla/ivona-go"
)

// Manager ...
type Manager interface {
	Listen()
}

type manager struct {
	tgClient    tg.Client
	ivonaClient *ivona.Ivona
}

// NewManager ...
func NewManager(tgClient tg.Client, ivonaClient *ivona.Ivona) Manager {
	return &manager{
		tgClient:    tgClient,
		ivonaClient: ivonaClient,
	}
}

// Listen ...
func (m *manager) Listen() {
	log.Println("Listening...")

	msgChan, _ := m.tgClient.GetMessagesChan()
	for msg := range msgChan {
		go m.handleMsg(msg)
	}
}

// handleMsg ...
func (m *manager) handleMsg(msg *tg.Message) {
	// TODO Валидация текста сообщения.
	log.Println("=====")
	log.Println(msg)
	log.Println(msg.Chat)

	options := ivona.NewSpeechOptions(msg.Text)
	options.Voice.Language = "ru-RU"
	options.Voice.Name = "Maxim"
	options.Voice.Gender = "Male"

	r, err := m.ivonaClient.CreateSpeech(options)
	if err != nil {
		log.Println(err)
	}

	m.tgClient.SendAudio(msg.Chat.ID, r.Audio)
}
