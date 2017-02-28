package app

import (
	"errors"
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
	ivonaClient ivonaInterface
	logger      logger
}

// NewManager ...
func NewManager(tgClient tg.Client, ivonaClient ivonaInterface) Manager {
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
func (m *manager) handleMsg(msg *tg.Message) error {
	if msg == nil {
		return errors.New("Nil pointer on msg")
	}

	options := m.buildIvonaOptions(msg.Text)

	r, err := m.ivonaClient.CreateSpeech(options)
	if err != nil {
		m.logger.Println(err)
		return err
	}

	m.tgClient.SendAudio(msg.Chat.ID, r.Audio)
	return nil
}

func (m *manager) buildIvonaOptions(text string) ivona.SpeechOptions {
	options := ivona.NewSpeechOptions(text)
	options.Voice.Language = "ru-RU"
	options.Voice.Name = "Maxim"
	options.Voice.Gender = "Male"
	return options
}
