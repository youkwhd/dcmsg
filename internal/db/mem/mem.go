package mem

import (
	"R2/internal/message"
)

var Messages = make(map[message.MessageID]message.Message)

func SaveMessage(channelID string, messageID string, role string, emoji string) {
	msg, found := Messages[message.MessageID(messageID)]
	if !found {
		msg = message.NewMessage(channelID)
	}

	msg.AddReaction(emoji, role)
	Messages[message.MessageID(messageID)] = msg
}

func GetMessage(messageID string) (message.Message, bool) {
	msg, found := Messages[message.MessageID(messageID)]
	return msg, found
}
