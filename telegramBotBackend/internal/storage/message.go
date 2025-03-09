package storage

// Тип сообщения. Пользовательский или бот
type MsgType string

const (
	USER MsgType = "USER"
	BOT  MsgType = "BOT"
)

type Message struct {
	Id      int
	text    string
	msgType MsgType
}

func NewMessage(id int, text string, msgType MsgType) *Message {
	return &Message{
		Id:      id,
		text:    text,
		msgType: msgType,
	}
}
