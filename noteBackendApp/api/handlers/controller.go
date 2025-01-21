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

type NoteController struct {
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

func (controller *NoteController) SaveNote(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	decoder := json.NewDecoder(context.Request.Body)
	var requestEntity noteApiDTO.CreateNoteRequest
	err := decoder.Decode(&requestEntity)
	if err != nil {
		var responseError = getErrorDto(err.Error(), http.StatusBadRequest, context)
		context.Status(http.StatusBadRequest)
		json.NewEncoder(context.Writer).Encode(responseError)
	}
	entity := controller.db.Save(noteDomain.NoteEntity{
		Id:   uuid.New().String(),
		Name: *requestEntity.Name,
		Link: *requestEntity.Link,
	})
	var response = noteApiDTO.NoteResponse{Id: &entity.Id, Name: &entity.Name, Link: &entity.Link}
	json.NewEncoder(context.Writer).Encode(response)
}

func (controller *NoteController) DeleteNoteById(context *gin.Context, id string) {
	context.Header("Content-Type", "application/json")
	controller.db.DeleteById(id)
	context.Status(http.StatusNoContent)
}

func (controller *NoteController) GetNoteById(context *gin.Context, id string) {
	context.Header("Content-Type", "application/json")
	obj, err := controller.db.GetById(id)
	if err != nil {
		var responseError = getErrorDto(err.Error(), http.StatusNotFound, context)
		context.Status(http.StatusNotFound)
		json.NewEncoder(context.Writer).Encode(responseError)
	} else {
		var response = noteApiDTO.NoteResponse{Id: &obj.Id, Name: &obj.Name, Link: &obj.Link}
		json.NewEncoder(context.Writer).Encode(response)
	}
}

func (controller *NoteController) GetAllNotes(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	allNotes := controller.db.GetAll()
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
