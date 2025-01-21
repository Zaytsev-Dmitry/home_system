package noteHandlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	noteApiDTO "noteBackendApp/api/openapi"
	noteDao "noteBackendApp/internal/dao"
	noteDomain "noteBackendApp/internal/domain"
	"time"
)

type NoteApi struct {
	db *noteDao.InMemoryNoteRepository
}

func getErrorDto(err string, status int, context *gin.Context) noteApiDTO.ErrorResponse {
	nowString := time.Now().String()
	return noteApiDTO.ErrorResponse{
		Timestamp: &nowString,
		Status:    &status,
		Error:     &err,
		Path:      &context.Request.URL.Path,
	}
}

func (api *NoteApi) SaveNote(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	decoder := json.NewDecoder(context.Request.Body)
	var requestEntity noteApiDTO.CreateNoteRequest
	err := decoder.Decode(&requestEntity)
	if err != nil {
		var responseError = getErrorDto(err.Error(), http.StatusBadRequest, context)
		context.Status(http.StatusBadRequest)
		json.NewEncoder(context.Writer).Encode(responseError)
	}
	entity := api.db.Save(noteDomain.NoteEntity{
		Id:   uuid.New().String(),
		Name: *requestEntity.Name,
		Link: *requestEntity.Link,
	})
	var response = noteApiDTO.NoteResponse{Id: &entity.Id, Name: &entity.Name, Link: &entity.Link}
	json.NewEncoder(context.Writer).Encode(response)
}

func (api *NoteApi) DeleteNoteById(context *gin.Context, id string) {
	context.Header("Content-Type", "application/json")
	api.db.DeleteById(id)
	context.Status(http.StatusNoContent)
}

func (api *NoteApi) GetNoteById(context *gin.Context, id string) {
	context.Header("Content-Type", "application/json")
	obj, err := api.db.GetById(id)
	if err != nil {
		var responseError = getErrorDto(err.Error(), http.StatusNotFound, context)
		context.Status(http.StatusNotFound)
		json.NewEncoder(context.Writer).Encode(responseError)
	} else {
		var response = noteApiDTO.NoteResponse{Id: &obj.Id, Name: &obj.Name, Link: &obj.Link}
		json.NewEncoder(context.Writer).Encode(response)
	}
}

func (api *NoteApi) GetAllNotes(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	allNotes := api.db.GetAll()
	var result = make([]noteApiDTO.NoteResponse, len(allNotes))
	for i, value := range allNotes {
		var response = noteApiDTO.NoteResponse{
			Id:   &value.Id,
			Link: &value.Link,
			Name: &value.Name,
		}
		result[i] = response
	}
	json.NewEncoder(context.Writer).Encode(result)
}

func NewNoteApi(db *noteDao.InMemoryNoteRepository) *NoteApi {
	return &NoteApi{db: db}
}
