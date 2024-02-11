package message

// Listened discord messages

type MessageID string
type Emoji string
type Role string

type Message struct {
	ChannelID string
	Reactions map[Emoji]Role
}

func NewMessage(channelID string) Message {
	return Message{
		ChannelID: channelID,
		Reactions: make(map[Emoji]Role),
	}
}
