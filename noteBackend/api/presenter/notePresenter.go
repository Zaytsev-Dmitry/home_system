package notePresenter

import (
	"fmt"
	generatedApi "noteBackendApp/api/spec"
	noteDomain "noteBackendApp/internal/domain"
)

type Presenter struct {
}

func (presenter *Presenter) ToNoteResponse(entity noteDomain.Note) generatedApi.NoteResponse {
	return generatedApi.NoteResponse{
		Name:        &entity.Name,
		Link:        &entity.Link,
		Description: &entity.Description,
		AccountId:   &entity.AccountId,
		TgId:        &entity.TelegramId,
	}
}

func (presenter *Presenter) ToEntity(requestEntity *generatedApi.CreateNoteRequest) noteDomain.Note {
	return noteDomain.Note{
		Name:        *requestEntity.Name,
		Description: *requestEntity.Description,
		Link:        *requestEntity.Link,
		TelegramId:  *requestEntity.TgId,
	}
}

func (presenter *Presenter) ToListNoteResponse(entities []noteDomain.Note) generatedApi.NoteResponseList {
	var result = make([]generatedApi.NoteResponse, len(entities))
	for i, value := range entities {
		id := fmt.Sprint(value.ID)
		var response = generatedApi.NoteResponse{
			AccountId:   &value.AccountId,
			TgId:        &value.TelegramId,
			Link:        &value.Link,
			Name:        &value.Name,
			Description: &value.Description,
			Id:          &id,
		}
		result[i] = response
	}

	return generatedApi.NoteResponseList{
		Objects: &result,
	}
}
