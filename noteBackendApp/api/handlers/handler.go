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

func (api *NoteApi) SaveNote(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	decoder := json.NewDecoder(context.Request.Body)
	var requestEntity noteApiDTO.CreateNoteRequest
	err := decoder.Decode(&requestEntity)
	if err != nil {
		errorMsg := err.Error()
		nowString := time.Now().String()
		httpStatus := http.StatusBadRequest
		var responseError = noteApiDTO.ErrorResponse{
			Timestamp: &nowString,
			Status:    &httpStatus,
			Error:     &errorMsg,
			Path:      &context.Request.URL.Path,
		}
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
		errorMsg := err.Error()
		nowString := time.Now().String()
		httpStatus := http.StatusNotFound
		var responseError = noteApiDTO.ErrorResponse{
			Timestamp: &nowString,
			Status:    &httpStatus,
			Error:     &errorMsg,
			Path:      &context.Request.URL.Path,
		}
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
