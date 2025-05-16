package service

type UserSession struct {
	Step    int
	Answers []string
}

type SessionStorage interface {
	Get(userID int64) *UserSession
	Set(userID int64, session *UserSession)
	Delete(userID int64)
	Reset(userID int64, totalQuestions int)
}

type InMemorySessionStorage struct {
	data map[int64]*UserSession
}

func NewInMemorySessionStorage() *InMemorySessionStorage {
	return &InMemorySessionStorage{
		data: make(map[int64]*UserSession),
	}
}

func (s InMemorySessionStorage) Get(userID int64) *UserSession {
	return s.data[userID]
}

func (s InMemorySessionStorage) Set(userID int64, session *UserSession) {
	s.data[userID] = session
}

func (s InMemorySessionStorage) Delete(userID int64) {
	delete(s.data, userID)
}

func (s InMemorySessionStorage) Reset(userID int64, totalQuestions int) {
	s.data[userID] = &UserSession{
		Step:    1,
		Answers: make([]string, totalQuestions),
	}
}
