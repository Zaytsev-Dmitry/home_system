package storage

// Тип сообщения. Пользовательский или бот
type msgType string

const (
	USER msgType = "USER"
	BOT  msgType = "BOT"
)

type Message struct {
	Id      int
	text    string
	msgType msgType
	index   int
}

func NewMessage(id int, text string, index int, msgType msgType) *Message {
	return &Message{
		Id:      id,
		text:    text,
		msgType: msgType,
		index:   index,
	}
}
