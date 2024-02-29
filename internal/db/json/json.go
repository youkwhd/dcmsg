package json

// Might be a painfully slow I/O operation
// but this is what we got for now.

import (
    "R2/internal/message"
    "encoding/json"
    "os"
)

func SaveMessage(channelID string, messageID string, role string, emoji string) {
    messages := GetAllMessages()
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

func GetMessage(messageID string) (message.Message, bool) {
    messages := GetAllMessages()
    msg, found := messages[message.MessageID(messageID)]
    return msg, found
}

// TODO: Cachable, maybe don't write it just yet
func GetAllMessages() map[message.MessageID]message.Message {
    bytes, _ := os.ReadFile("data/db.json")

    messages := make(map[message.MessageID]message.Message)
    json.Unmarshal(bytes, &messages)

    return messages
}

func replaceJsonData(messages map[message.MessageID]message.Message) {
    bytes, _ := json.Marshal(messages)

    os.Mkdir("data", os.ModePerm)
    os.WriteFile("data/db.json", bytes, 0666)
}

// Removes the message if found
func RemoveMessage(messageID string) {
    messages := GetAllMessages()

    _, found := messages[message.MessageID(messageID)]
    if !found {
        return;
    }

    delete(messages, message.MessageID(messageID));
    replaceJsonData(messages)
}
