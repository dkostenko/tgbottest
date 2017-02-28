package tg

import "encoding/json"

// Message описывает сообщение Telegram.
type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
	Chat      *Chat  `json:"chat"`
}

// Chat описывает чат в системе Telegram.
type Chat struct {
	ID int `json:"id"`
}

// Update https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message"`
}

// Response -
type Response struct {
	Ok        bool            `json:"ok"`
	Result    json.RawMessage `json:"result"`
	ErrorCode int             `json:"error_code"`
}
