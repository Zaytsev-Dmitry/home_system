package notePresenter

import (
	noteApiDTO "noteBackendApp/api/docs"
	noteDomain "noteBackendApp/internal/domain"
)

type Presenter struct {
}

func (presenter *Presenter) ToNoteResponse(entity noteDomain.TelegramAccount) noteApiDTO.NoteResponse {
	return noteApiDTO.NoteResponse{
		Name:        &entity.Name,
		Link:        entity.Link,
		Description: entity.Description,
		AccountId:   &entity.AccountId,
	}
}

func (presenter *Presenter) ToEntity(requestEntity *noteApiDTO.CreateNoteRequest) noteDomain.TelegramAccount {
	return noteDomain.TelegramAccount{
		Name:        *requestEntity.Name,
		AccountId:   *requestEntity.AccountId,
		Description: requestEntity.Description,
		Link:        requestEntity.Link,
	}
}

func (presenter *Presenter) ToListNoteResponse(entities []noteDomain.TelegramAccount) []noteApiDTO.NoteResponse {
	var result = make([]noteApiDTO.NoteResponse, len(entities))
	for i, value := range entities {
		var response = noteApiDTO.NoteResponse{
			Link: value.Link,
			Name: &value.Name,
		}
		result[i] = response
	}
	return result
}
