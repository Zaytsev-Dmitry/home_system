package noteApi

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
	noteDao "noteBackendApp/internal/dao"
	noteDomain "noteBackendApp/internal/domain"
)

// TODO добавить валидацию
func getById(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	obj, err := noteDao.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		var response = noteDomain.NoteResponse{Id: obj.Id, Name: obj.Name, Link: obj.Link}
		json.NewEncoder(w).Encode(response)
	}
}

func save(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var requestEntity noteDomain.CreateNoteRequest
	err := decoder.Decode(&requestEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	entity := noteDao.Save(noteDomain.NoteEntity{Id: uuid.New().String(), Name: requestEntity.Name, Link: requestEntity.Link})
	var response = noteDomain.NoteResponse{Id: entity.Id, Name: entity.Name, Link: entity.Link}

	json.NewEncoder(w).Encode(response)
}

func Init() *httprouter.Router {
	router := httprouter.New()
	router.GET("/note/:id", getById)
	router.POST("/note", save)
	return router
}
