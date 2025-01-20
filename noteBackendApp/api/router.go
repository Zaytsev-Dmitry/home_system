package noteApi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	noteDao "noteBackendApp/internal/dao"
	noteDomain "noteBackendApp/internal/domain"
	"time"
)

func getById(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	id := context.Params.ByName("id")
	obj, err := noteDao.GetById(id)
	if err != nil {
		var responseError = noteDomain.ResponseError{
			Timestamp:    time.Now().String(),
			Status:       http.StatusNotFound,
			BusinessCode: "-",
			Error:        err.Error(),
			Path:         context.Request.URL.Path,
		}
		json.NewEncoder(context.Writer).Encode(responseError)
	} else {
		var response = noteDomain.NoteResponse{Id: obj.Id, Name: obj.Name, Link: obj.Link}
		json.NewEncoder(context.Writer).Encode(response)
	}
}

func save(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	decoder := json.NewDecoder(context.Request.Body)
	var requestEntity noteDomain.CreateNoteRequest
	err := decoder.Decode(&requestEntity)
	if err != nil {
		var responseError = noteDomain.ResponseError{
			Timestamp:    time.Now().String(),
			Status:       http.StatusBadRequest,
			BusinessCode: "-",
			Error:        err.Error(),
			Path:         context.Request.URL.Path,
		}
		json.NewEncoder(context.Writer).Encode(responseError)
	}
	entity := noteDao.Save(noteDomain.NoteEntity{Id: uuid.New().String(), Name: requestEntity.Name, Link: requestEntity.Link})
	var response = noteDomain.NoteResponse{Id: entity.Id, Name: entity.Name, Link: entity.Link}
	json.NewEncoder(context.Writer).Encode(response)
}

func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/note/:id", getById)
	router.POST("/note", save)
	return router
}
