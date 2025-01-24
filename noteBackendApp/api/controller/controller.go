package noteControllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	noteApiDTO "noteBackendApp/api/docs"
	notePresenter "noteBackendApp/api/presenter"
	noteDao "noteBackendApp/internal/dao/impl"
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
	var requestEntity noteApiDTO.CreateNoteRequest
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
	obj, err := controller.GetUseCase.GetNoteByAccountId(accountId)
	if err != nil {
		noteUtilities.SetResponseError(err, context, http.StatusNotFound)
	} else {
		noteUtilities.SetResponse(
			noteApiDTO.NoteResponse{Name: &obj.Name, Link: obj.Link},
			context,
		)
	}
}

func Create(db *noteDao.InMemoryNoteRepository) *NoteController {
	return &NoteController{
		SaveUseCase:   &noteUseCase.SaveNoteUseCase{Db: db},
		DeleteUseCase: &noteUseCase.DeleteNoteUseCase{Db: db},
		GetUseCase:    &noteUseCase.GetNoteUseCase{Db: db},
		presenter:     &notePresenter.Presenter{},
	}
}
