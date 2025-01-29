package noteControllers

import (
	noteSpec "github.com/Zaytsev-Dmitry/home_system_open_api/noteServerBackend"
	"github.com/gin-gonic/gin"
	"net/http"
	notePresenter "noteBackendApp/api/presenter"
	noteInterface "noteBackendApp/internal/dao/interface"
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
	var requestEntity noteSpec.CreateNoteRequest
	noteUtilities.CatchMarshallErr(c.BindJSON(&requestEntity), c)
	entity := controller.SaveUseCase.Save(controller.presenter.ToEntity(&requestEntity))
	noteUtilities.SetResponse(
		controller.presenter.ToNoteResponse(entity),
		c,
	)
}

func (controller *NoteController) DeleteNotesByAccountId(context *gin.Context, accountId int) {
	controller.DeleteUseCase.DeleteNoteByAccountId(accountId)
	context.Status(http.StatusNoContent)
}

func (controller *NoteController) GetNotesByAccountId(context *gin.Context, accountId int) {
	obj := controller.GetUseCase.GetNoteByAccountId(accountId)
	noteUtilities.SetResponse(
		controller.presenter.ToListNoteResponse(obj),
		context,
	)
}

func Create(dao noteInterface.NoteDao) *NoteController {
	return &NoteController{
		SaveUseCase:   &noteUseCase.SaveNoteUseCase{DAO: dao},
		DeleteUseCase: &noteUseCase.DeleteNoteUseCase{DAO: dao},
		GetUseCase:    &noteUseCase.GetNoteUseCase{DAO: dao},
		presenter:     &notePresenter.Presenter{},
	}
}
