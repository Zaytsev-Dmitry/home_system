package noteHandlers

import (
	"github.com/gin-gonic/gin"
	"noteBackendApp/api/controller"
	"noteBackendApp/internal/app/ports/out/dao"
)

type NoteBackendApi struct {
	controller *controller.NoteController
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

func NewNoteBackendApi(dao *dao.NoteDao) *NoteBackendApi {
	return &NoteBackendApi{controller: controller.Create(dao)}
}
