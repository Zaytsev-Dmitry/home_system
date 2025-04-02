package notePresenter

import (
	"fmt"
	openapi "noteBackendApp/api/http"
	noteDomain "noteBackendApp/internal/domain"
)

type Presenter struct {
}

func (presenter *Presenter) ToNoteResponse(entity noteDomain.Note) openapi.NoteResponse {
	return openapi.NoteResponse{
		Name:        &entity.Name,
		Link:        &entity.Link,
		Description: &entity.Description,
		AccountId:   &entity.AccountId,
		TgId:        &entity.TelegramId,
	}
}

func (presenter *Presenter) ToEntity(requestEntity *openapi.CreateNoteRequest) noteDomain.Note {
	return noteDomain.Note{
		Name:        *requestEntity.Name,
		Description: *requestEntity.Description,
		Link:        *requestEntity.Link,
		TelegramId:  *requestEntity.TgId,
	}
}

func (presenter *Presenter) ToListNoteResponse(entities []noteDomain.Note) openapi.NoteResponseList {
	var result = make([]openapi.NoteResponse, len(entities))
	for i, value := range entities {
		id := fmt.Sprint(value.ID)
		var response = openapi.NoteResponse{
			AccountId:   &value.AccountId,
			TgId:        &value.TelegramId,
			Link:        &value.Link,
			Name:        &value.Name,
			Description: &value.Description,
			Id:          &id,
		}
		result[i] = response
	}

	return openapi.NoteResponseList{
		Objects: &result,
	}
}
