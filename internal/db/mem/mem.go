package mem

import (
	"R2/internal/message"
)

var Messages = make(map[message.MessageID]message.Message)

func SaveMessage(channelID string, messageID string, role string, emoji string) {
	Messages[message.MessageID(messageID)] = message.NewMessage(channelID)
	Messages[message.MessageID(messageID)].Reactions[message.Emoji(emoji)] = message.Role(role)
}

func GetMessage(messageID string) (message.Message, bool) {
	msg, found := Messages[message.MessageID(messageID)]
	return msg, found
}
