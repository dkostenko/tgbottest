package app

import (
	"errors"
	"testing"

	"github.com/dkostenko/tgbottest/tg"
	ivona "github.com/jpadilla/ivona-go"
)

type tgClient struct{}

func (b *tgClient) GetMessagesChan() (<-chan *tg.Message, error) { return nil, nil }
func (b *tgClient) SendAudio(chatID int, audio []byte)           {}

type iClient struct{}

func (c *iClient) CreateSpeech(ivona.SpeechOptions) (*ivona.SpeechResponse, error) {
	return nil, errors.New("ERROR")
}

func TestHandleMsg(t *testing.T) {
	tgClient := &tgClient{}
	ivonaClient := &iClient{}

	m := &manager{tgClient: tgClient, ivonaClient: ivonaClient}

	err := m.handleMsg(nil)
	if err.Error() != "Nil pointer on msg" {
		t.Error("fail 1")
	}

	err = m.handleMsg(&tg.Message{})
	if err.Error() != "ERROR" {
		t.Error("fail 2")
	}
}
