package noteHandlers

import (
	"github.com/gin-gonic/gin"
	noteDao "noteBackendApp/internal/dao"
)

type NoteBackendApi struct {
	controller *NoteController
}

func (api *NoteBackendApi) SaveNote(context *gin.Context) {
	api.controller.SaveNote(context)
}

func (api *NoteBackendApi) DeleteNoteById(context *gin.Context, id string) {
	api.controller.DeleteNoteById(context, id)
}

func (api *NoteBackendApi) GetNoteById(context *gin.Context, id string) {
	api.controller.GetNoteById(context, id)
}

func (api *NoteBackendApi) GetAllNotes(context *gin.Context) {
	api.controller.GetAllNotes(context)
}

func NewNoteBackendApi(db *noteDao.InMemoryNoteRepository) *NoteBackendApi {
	return &NoteBackendApi{controller: &NoteController{db: db}}
}
