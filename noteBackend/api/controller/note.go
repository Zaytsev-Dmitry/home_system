package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	openapi "noteBackendApp/api/http"
	notePresenter "noteBackendApp/api/presenter"
	noteInterface "noteBackendApp/internal/dao"
	noteUseCase "noteBackendApp/internal/usecases"
	noteUtilities "noteBackendApp/pkg/utilities"
)

type NoteController struct {
	SaveUseCase   *noteUseCase.SaveNoteUseCase
	DeleteUseCase *noteUseCase.DeleteNoteUseCase
	GetUseCase    *noteUseCase.GetNoteUseCase
	presenter     *notePresenter.Presenter
}

func (controller *NoteController) SaveNote(c *gin.Context) {
	var requestEntity openapi.CreateNoteRequest
	noteUtilities.CatchMarshallErr(c.BindJSON(&requestEntity), c)
	entity := controller.SaveUseCase.Save(controller.presenter.ToEntity(&requestEntity))
	noteUtilities.SetResponse(
		controller.presenter.ToNoteResponse(entity),
		c,
	)
}

func (controller *NoteController) DeleteNotesByTgId(context *gin.Context, tgId int64) {
	controller.DeleteUseCase.DeleteNoteByTgId(tgId)
	context.Status(http.StatusNoContent)
}

func (controller *NoteController) GetNotesByTgId(context *gin.Context, tgId int64) {
	obj := controller.GetUseCase.GetNoteByTgId(tgId)
	noteUtilities.SetResponse(
		controller.presenter.ToListNoteResponse(obj),
		context,
	)
}

func Create(dao noteInterface.NoteDao) *NoteController {
	return &NoteController{
		SaveUseCase:   &noteUseCase.SaveNoteUseCase{Dao: dao},
		DeleteUseCase: &noteUseCase.DeleteNoteUseCase{Dao: dao},
		GetUseCase:    &noteUseCase.GetNoteUseCase{Dao: dao},
		presenter:     &notePresenter.Presenter{},
	}
}
