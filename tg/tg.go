package tg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client - менеджер общения с Telegram Bot API.
type Client interface {
	GetMessagesChan() (<-chan *Message, error)
}

type client struct {
	token     string
	lastUpdID int
}

// Message описывает сообщение Telegram.
type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
}

// Update https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message"`
}

// Response -
type Response struct {
	Ok     bool      `json:"ok"`
	Result []*Update `json:"result"`
}

// NewClient - конструктор менеджера общения с Telegram Bot API.
func NewClient(token string) Client {
	return &client{token: token}
}

func (c *client) makeRequest(uri string) (*Response, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetMessagesChan возвращает канал с сообщениями.
func (c *client) GetMessagesChan() (<-chan *Message, error) {
	msgChan := make(chan *Message)

	go func() {
		for {
			uri := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?timeout=100&offset=%d",
				c.token, c.lastUpdID+1)
			resp, _ := c.makeRequest(uri)
			for _, upd := range resp.Result {
				c.lastUpdID = upd.UpdateID
				msgChan <- upd.Message
			}
		}
	}()

	return msgChan, nil
}
