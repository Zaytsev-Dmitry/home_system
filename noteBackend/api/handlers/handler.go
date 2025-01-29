package noteHandlers

import (
	"github.com/gin-gonic/gin"
	noteController "noteBackendApp/api/controller"
	noteInterface "noteBackendApp/internal/dao/interface"
)

type NoteBackendApi struct {
	controller *noteController.NoteController
}

func (api *NoteBackendApi) SaveNote(c *gin.Context) {
	api.controller.SaveNote(c)
}

func (api *NoteBackendApi) DeleteNotesByAccountId(c *gin.Context, accountId int) {
	api.controller.DeleteNotesByAccountId(c, accountId)
}

func (api *NoteBackendApi) GetNotesByAccountId(c *gin.Context, accountId int) {
	api.controller.GetNotesByAccountId(c, accountId)
}

func NewNoteBackendApi(dao noteInterface.NoteDao) *NoteBackendApi {
	return &NoteBackendApi{controller: noteController.Create(dao)}
}
