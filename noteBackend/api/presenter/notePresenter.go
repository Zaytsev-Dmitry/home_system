package notePresenter

import (
	noteSpec "github.com/Zaytsev-Dmitry/home_system_open_api/noteServerBackend"
	noteDomain "noteBackendApp/internal/domain"
)

type Presenter struct {
}

func (presenter *Presenter) ToNoteResponse(entity noteDomain.Note) noteSpec.NoteResponse {
	return noteSpec.NoteResponse{
		Name:        &entity.Name,
		Link:        entity.Link,
		Description: entity.Description,
		AccountId:   &entity.AccountId,
		TgId:        &entity.TelegramId,
	}
}

func (presenter *Presenter) ToEntity(requestEntity *noteSpec.CreateNoteRequest) noteDomain.Note {
	return noteDomain.Note{
		Name:        *requestEntity.Name,
		Description: requestEntity.Description,
		Link:        requestEntity.Link,
		TelegramId:  *requestEntity.TgId,
	}
}

func (presenter *Presenter) ToListNoteResponse(entities []noteDomain.Note) noteSpec.NoteResponseList {
	var result = make([]noteSpec.NoteResponse, len(entities))
	for i, value := range entities {
		var response = noteSpec.NoteResponse{
			AccountId:   &value.AccountId,
			TgId:        &value.TelegramId,
			Link:        value.Link,
			Name:        &value.Name,
			Description: value.Description,
		}
		result[i] = response
	}

	return noteSpec.NoteResponseList{
		Objects: &result,
	}
}
