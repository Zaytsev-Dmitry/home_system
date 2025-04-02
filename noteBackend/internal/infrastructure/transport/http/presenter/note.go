package presenter

import (
	"github.com/gin-gonic/gin"
	openapi "noteBackendApp/api/http"
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao/dto"
	"time"
)

type Presenter struct {
}

func (p *Presenter) PresentToReq(req openapi.CreateNoteRequest) dto.CreateNoteDto {
	return dto.CreateNoteDto{
		TgId:        *req.TgId,
		Description: *req.Description,
		Link:        *req.Link,
		Name:        *req.Name,
	}
}

func (p *Presenter) toNoteResponse(note *domain.Note) *openapi.NoteResponse {
	return &openapi.NoteResponse{
		Description: &note.Description,
		Id:          &note.ID,
		Link:        &note.Link,
		Name:        &note.Name,
	}
}

func (p *Presenter) toMetaData(context *gin.Context) *openapi.MetaData {
	nowString := time.Now().String()
	return &openapi.MetaData{
		Path:      &context.Request.URL.Path,
		Timestamp: &nowString,
	}
}

func (p *Presenter) PresentToResp(note *domain.Note, context *gin.Context) openapi.SingleNoteBackendResponse {
	return openapi.SingleNoteBackendResponse{
		Payload: p.toNoteResponse(note),
		Meta:    p.toMetaData(context),
	}
}
func (p *Presenter) PresentToListResp(notes []*domain.Note, context *gin.Context) openapi.ListNoteBackendResponse {
	noteResp := make([]openapi.NoteResponse, 0)
	for _, note := range notes {
		noteResp = append(noteResp, *p.toNoteResponse(note))
	}
	return openapi.ListNoteBackendResponse{
		Payload: &noteResp,
		Meta:    p.toMetaData(context),
	}
}
