package notePresenter

import (
	"github.com/google/uuid"
	noteApiDTO "noteBackendApp/api/docs"
	noteDomain "noteBackendApp/internal/domain"
)

type Presenter struct {
}

func (presenter *Presenter) ToNoteResponse(entity noteDomain.NoteEntity) noteApiDTO.NoteResponse {
	return noteApiDTO.NoteResponse{Id: &entity.Id, Name: &entity.Name, Link: &entity.Link}
}

func (presenter *Presenter) ToEntity(requestEntity *noteApiDTO.CreateNoteRequest) noteDomain.NoteEntity {
	return noteDomain.NoteEntity{
		Id:   uuid.New().String(),
		Name: *requestEntity.Name,
		Link: *requestEntity.Link,
	}
}

func (presenter *Presenter) ToListNoteResponse(entities []noteDomain.NoteEntity) []noteApiDTO.NoteResponse {
	var result = make([]noteApiDTO.NoteResponse, len(entities))
	for i, value := range entities {
		var response = noteApiDTO.NoteResponse{
			Id:   &value.Id,
			Link: &value.Link,
			Name: &value.Name,
		}
		result[i] = response
	}
	return result
}
