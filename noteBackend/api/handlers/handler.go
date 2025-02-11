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

func (api *NoteBackendApi) DeleteNotesByTgId(c *gin.Context, tgId int) {
	api.controller.DeleteNotesByTgId(c, int64(tgId))
}

func (api *NoteBackendApi) GetNotesByTgId(c *gin.Context, tgId int64) {
	api.controller.GetNotesByTgId(c, tgId)
}

func NewNoteBackendApi(dao noteInterface.NoteDao) *NoteBackendApi {
	return &NoteBackendApi{controller: noteController.Create(dao)}
}
