package message

// Listened discord messages

type MessageID string
type Emoji string
type Role string

type Message struct {
    ChannelID string             `json:"channel_id"`
    Reactions map[Emoji]Role     `json:"reactions"`
}

func NewMessage(channelID string) Message {
    return Message{
        ChannelID: channelID,
        Reactions: make(map[Emoji]Role),
    }
}

func (msg *Message) AddReaction(emoji string, role string) {
    msg.Reactions[Emoji(emoji)] = Role(role)
}
