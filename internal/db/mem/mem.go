package mem

type MessageID string
type Emoji string
type Role string

type RoleReactionMessage struct {
	ChannelID string
	Reactions map[Emoji]Role
}

func NewRoleReactionMessage(channelID string) RoleReactionMessage {
	return RoleReactionMessage{
		ChannelID: channelID,
		Reactions: make(map[Emoji]Role),
	}
}

var Messages = make(map[MessageID]RoleReactionMessage)
