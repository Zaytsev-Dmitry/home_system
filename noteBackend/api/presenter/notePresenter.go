package notePresenter

import (
	noteSpec "github.com/Zaytsev-Dmitry/home_system_open_api/noteServerBackend"
	noteDomain "noteBackendApp/internal/domain"
)

type Presenter struct {
}

func (presenter *Presenter) ToNoteResponse(entity noteDomain.TelegramAccount) noteSpec.NoteResponse {
	return noteSpec.NoteResponse{
		Name:        &entity.Name,
		Link:        entity.Link,
		Description: entity.Description,
		AccountId:   &entity.AccountId,
	}
}

func (presenter *Presenter) ToEntity(requestEntity *noteSpec.CreateNoteRequest) noteDomain.TelegramAccount {
	return noteDomain.TelegramAccount{
		Name:        *requestEntity.Name,
		AccountId:   *requestEntity.AccountId,
		Description: requestEntity.Description,
		Link:        requestEntity.Link,
	}
}

func (presenter *Presenter) ToListNoteResponse(entities []noteDomain.TelegramAccount) []noteSpec.NoteResponse {
	var result = make([]noteSpec.NoteResponse, len(entities))
	for i, value := range entities {
		var response = noteSpec.NoteResponse{
			Link: value.Link,
			Name: &value.Name,
		}
		result[i] = response
	}
	return result
}
