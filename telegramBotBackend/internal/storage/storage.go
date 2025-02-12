package storage

var messages = make(map[int64][]Message)

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Add(tgID int64, message Message) {
	messages[tgID] = append(messages[tgID], message)
}

func (s *Storage) Get(tgID int64) []Message {
	return messages[tgID]
}

func (s *Storage) ClearAll(tgID int64) {
	delete(messages, tgID)
}
