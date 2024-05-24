package mem

import (
    "rrolls/internal/message"
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

func GetAllMessages() map[message.MessageID]message.Message {
    return Messages
}

// Removes the message if found
func RemoveMessage(messageID string) {
    _, found := Messages[message.MessageID(messageID)]
    if !found {
        return;
    }

    delete(Messages, message.MessageID(messageID))
}
