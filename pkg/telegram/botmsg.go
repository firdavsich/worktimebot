package telegram

import (
	"fmt"
	"net/http"
)

// Bot struct
type Bot struct {
	Token  string
	Silent bool
}

// Send massage to chat
func (b *Bot) Send(chatID int64, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?disable_notification=%t&chat_id=%d&text=%s", b.Token, b.Silent, chatID, text)
	_, err := http.Get(url)
	if err != nil {
		return err
	}

	return nil
}
