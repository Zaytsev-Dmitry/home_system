package services

import (
	"database/sql"
	"errors"
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao/dto"
	"noteBackendApp/internal/app/ports/out/dao/repository/note"
	customErr "noteBackendApp/pkg/errors"
)

type SaveNoteUCaseImpl struct {
	NoteRepo note.NoteRepository
}

func (s *SaveNoteUCaseImpl) Save(dto dto.CreateNoteDto) (*domain.Note, error) {
	// Проверяем существование заметки
	entity, err := s.NoteRepo.ExistByName(dto.Name)
	if err == nil {
		return entity, nil // Если запись найдена, возвращаем её сразу
	}

	// Если ошибка означает, что запись не найдена, создаем новую
	if errors.Is(err.WrapErr, sql.ErrNoRows) {
		return s.createNewNote(dto)
	}

	// В случае других ошибок возвращаем обработанное сообщение
	return nil, customErr.UpdateErrorText(err.Action, err.WrapErr)
}

func (s *SaveNoteUCaseImpl) createNewNote(dto dto.CreateNoteDto) (*domain.Note, error) {
	entity, err := s.NoteRepo.Save(domain.Note{
		Description: dto.Description,
		Link:        dto.Link,
		Name:        dto.Name,
		TelegramId:  dto.TgId,
	})

	if err != nil {
		return nil, customErr.UpdateErrorText(err.Action, err.WrapErr)
	}
	return entity, nil
}
