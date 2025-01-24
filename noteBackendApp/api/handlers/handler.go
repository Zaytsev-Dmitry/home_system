package noteHandlers

import (
	"github.com/gin-gonic/gin"
	noteController "noteBackendApp/api/controller"
	noteDao "noteBackendApp/internal/dao/impl"
)

type NoteBackendApi struct {
	controller *noteController.NoteController
}

func (api *NoteBackendApi) SaveNote(c *gin.Context) {
	api.controller.SaveNote(c)
}

func (api *NoteBackendApi) DeleteNoteByAccountId(c *gin.Context, accountId int) {
	api.controller.DeleteNoteByAccountId(c, accountId)
}

func (api *NoteBackendApi) GetNoteByAccountId(c *gin.Context, accountId int) {
	api.controller.GetNoteByAccountId(c, accountId)
}

func NewNoteBackendApi(db *noteDao.InMemoryNoteRepository) *NoteBackendApi {
	return &NoteBackendApi{controller: noteController.Create(db)}
}
