package app

import (
	"log"

	ivona "github.com/jpadilla/ivona-go"
)

type ivonaInterface interface {
	CreateSpeech(ivona.SpeechOptions) (*ivona.SpeechResponse, error)
}

type loggerInterface interface {
	Println(v ...interface{})
}

type logger struct{}

func (l *logger) Println(v ...interface{}) {
	log.Println(v)
}
