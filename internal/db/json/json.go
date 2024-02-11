package json

import (
	"R2/internal/message"
)

func SaveMessage(channelID string, messageID string, role string, emoji string) {

}

func GetMessage(messageID string) (message.Message, bool) {
	return message.Message{}, false
}
