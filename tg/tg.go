package tg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

// Client - менеджер общения с Telegram Bot API.
type Client interface {
	GetMessagesChan() (<-chan *Message, error)
	SendAudio(chatID int, audio []byte)
}

type client struct {
	token     string
	lastUpdID int
}

// NewClient - конструктор менеджера общения с Telegram Bot API.
func NewClient(token string) Client {
	return &client{token: token}
}

func (c *client) makeGETRequest(uri string) (*Response, error) {
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

func (c *client) makePOSTRequest(uri, fileKey, fileName string, file []byte,
	fields map[string]string) (*Response, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	buf := bytes.NewBuffer(file)

	fw, err := w.CreateFormFile(fileKey, fileName)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, buf); err != nil {
		log.Println(err)
		return nil, err
	}

	for k, v := range fields {
		w.WriteField(k, v)
	}

	w.Close()

	req, err := http.NewRequest("POST", uri, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := &Response{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetMessagesChan возвращает канал с сообщениями.
func (c *client) GetMessagesChan() (<-chan *Message, error) {
	msgChan := make(chan *Message)

	go func() {
		for {
			uri := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?timeout=100&offset=%d",
				c.token, c.lastUpdID+1)
			resp, _ := c.makeGETRequest(uri)

			var updates []*Update
			err := json.Unmarshal(resp.Result, &updates)
			if err != nil {
				panic(err)
			}

			for _, upd := range updates {
				c.lastUpdID = upd.UpdateID
				msgChan <- upd.Message
			}
		}
	}()

	return msgChan, nil
}

// SendAudio ...
func (c *client) SendAudio(chatID int, audio []byte) {
	uri := fmt.Sprintf("https://api.telegram.org/bot%s/sendAudio", c.token)

	resp, err := c.makePOSTRequest(uri, "audio", "audio.mp3", audio, map[string]string{
		"chat_id": strconv.Itoa(chatID),
	})
	if err != nil {
		log.Println(err)
		return
	}
	if !resp.Ok {
		err := fmt.Errorf("Tg error %d", resp.ErrorCode)
		log.Println(err)
		return
	}
}
