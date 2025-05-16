package registry

import (
	"botCoreService/internal/infrastructure/transport/tg/commands"
)

type InMemoryCommandRegistry struct {
	sessions map[int64]commands.BaseCommand
}

func NewCommandSessionManager() *InMemoryCommandRegistry {
	return &InMemoryCommandRegistry{
		sessions: make(map[int64]commands.BaseCommand),
	}
}

// Установить активную команду для пользователя
func (m *InMemoryCommandRegistry) Set(userID int64, cmd commands.BaseCommand) {
	m.sessions[userID] = cmd
}

// Получить текущую команду пользователя
func (m *InMemoryCommandRegistry) Get(userID int64) (commands.BaseCommand, bool) {
	cmd, ok := m.sessions[userID]
	return cmd, ok
}

// Удаляет (пользак закончил исполнение команды) текущую команду пользователя
func (m *InMemoryCommandRegistry) Delete(userID int64) {
	delete(m.sessions, userID)
}
