package json

// Might be a painfully slow I/O operation
// but this is what we got for now.

import (
    "R2/internal/message"
    "encoding/json"
    "os"
)

func SaveMessage(channelID string, messageID string, role string, emoji string) {
    messages := getMessages()
    msg, found := messages[message.MessageID(messageID)]
    if !found {
        msg = message.NewMessage(channelID)
    }

    msg.AddReaction(emoji, role)
    messages[message.MessageID(messageID)] = msg

    bytes, _ := json.Marshal(messages)

    os.Mkdir("data", os.ModePerm)
    os.WriteFile("data/db.json", bytes, 0666)
}

// TODO: Cachable, maybe don't write it just yet
func getMessages() map[message.MessageID]message.Message {
    bytes, _ := os.ReadFile("data/db.json")

    messages := make(map[message.MessageID]message.Message)
    json.Unmarshal(bytes, &messages)

    return messages
}

func GetMessage(messageID string) (message.Message, bool) {
    messages := getMessages()
    msg, found := messages[message.MessageID(messageID)]
    return msg, found
}
