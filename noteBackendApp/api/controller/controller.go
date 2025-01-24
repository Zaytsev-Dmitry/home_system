package noteControllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	noteApiDTO "noteBackendApp/api/docs"
	notePresenter "noteBackendApp/api/presenter"
	noteDao "noteBackendApp/internal/dao"
	noteUseCase "noteBackendApp/internal/usecases"
	noteUtilities "noteBackendApp/pkg/utilities"
)

type NoteController struct {
	SaveUseCase   *noteUseCase.SaveNoteUseCase
	DeleteUseCase *noteUseCase.DeleteNoteUseCase
	GetUseCase    *noteUseCase.GetNoteUseCase
	presenter     *notePresenter.Presenter
}

func (controller *NoteController) SaveNote(context *gin.Context) {
	var requestEntity noteApiDTO.CreateNoteRequest
	noteUtilities.CatchMarshallErr(context.BindJSON(&requestEntity), context)
	entity := controller.SaveUseCase.Save(controller.presenter.ToEntity(&requestEntity))
	noteUtilities.SetResponse(
		controller.presenter.ToNoteResponse(entity),
		context,
	)
}

func (controller *NoteController) DeleteNoteByAccountId(context *gin.Context, accountId int) {
	controller.DeleteUseCase.DeleteNoteByAccountId(accountId)
	context.Status(http.StatusNoContent)
}

func (controller *NoteController) GetNoteByAccountId(context *gin.Context, accountId int) {
	obj, err := controller.GetUseCase.GetNoteByAccountId(accountId)
	if err != nil {
		noteUtilities.SetResponseError(err, context, http.StatusNotFound)
	} else {
		noteUtilities.SetResponse(
			noteApiDTO.NoteResponse{Id: &obj.Id, Name: &obj.Name, Link: &obj.Link},
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
