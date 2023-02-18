package api

import (
	"fmt"
	"net/http"
	"net/url"
)

type Bot struct {
	Client string
}

func NewBot(token string) *Bot {
	return &Bot{
		Client: fmt.Sprintf("https://api.telegram.org/bot%s", token),
	}
}

func (b *Bot) SendMessage(message, uuid string) error {
	data := url.Values{
		"chat_id": {uuid},
		"text":    {"Web3000: " + message},
	}
	_, err := http.PostForm(fmt.Sprintf("%s/sendMessage", b.Client), data)
	if err != nil {
		return err
	}
	return nil
}
