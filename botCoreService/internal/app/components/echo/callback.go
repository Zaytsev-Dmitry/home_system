package echo

import (
	"botCoreService/internal/app/service"
	"botCoreService/pkg/telegram"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

const (
	CONFIRM_YES = "yes"
	CONFIRM_NO  = "no"
)

// Основной обработчик callback'ов компонента Echo
func (e *Echo) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := telegram.GetUserId(update)
	session := getOrCreateSession(e, userID, len(e.Question))

	switch {
	case update.CallbackQuery != nil:
		handleConfirmation(ctx, b, update, userID, e, session)

	case update.Message != nil:
		handleAnswerInput(ctx, b, update, userID, e, session)
	}
}

func getOrCreateSession(e *Echo, userID int64, totalQuestions int) *service.UserSession {
	session := e.sessionStorage.Get(userID)
	if session != nil {
		return session
	}
	session = &service.UserSession{
		Step:    0,
		Answers: make([]string, totalQuestions),
	}
	e.sessionStorage.Set(userID, session)
	return session
}

func handleAnswerInput(ctx context.Context, b *bot.Bot, update *models.Update, userID int64, e *Echo, session *service.UserSession) {
	if session.Step > 0 && session.Step <= len(session.Answers) {
		session.Answers[session.Step-1] = update.Message.Text
	}
	if session.Step < len(e.Question) {
		sendNextQuestion(ctx, b, userID, e, session)
	} else {
		sendConfirmationRequest(ctx, b, userID, e)
	}
}

func sendNextQuestion(ctx context.Context, b *bot.Bot, userID int64, e *Echo, session *service.UserSession) {
	q := e.Question[session.Step]
	session.Step++

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: userID,
		Text:   q.Content,
	})
}

func handleConfirmation(ctx context.Context, b *bot.Bot, update *models.Update, userID int64, e *Echo, session *service.UserSession) {
	cmd := strings.TrimPrefix(update.CallbackQuery.Data, e.prefix)

	switch cmd {
	case CONFIRM_YES:
		result := EchoResult{
			ChatId:   userID,
			Question: mergeAnswers(e.Question, session.Answers),
		}
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text:   e.text.ConfirmText,
		})
		e.proceedResult(result)
		e.CommandRegistry.Delete(userID)
		e.sessionStorage.Delete(userID)
	case CONFIRM_NO:
		resetSession(e, userID, len(e.Question))
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text:   "Ну хорошо, давай заново: " + e.Question[0].Content,
		})
	default:
		sendNextQuestion(ctx, b, userID, e, session)
	}
}
func resetSession(e *Echo, userID int64, totalQuestions int) {
	e.sessionStorage.Reset(userID, totalQuestions)
}

func sendConfirmationRequest(ctx context.Context, b *bot.Bot, userID int64, e *Echo) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      userID,
		Text:        formatConfirmationText(e.text.ConfirmText, mergeAnswers(e.Question, e.sessionStorage.Get(userID).Answers)),
		ReplyMarkup: e.buildDefaultConfirmKeyboard(),
		ParseMode:   models.ParseModeHTML,
	})
}

func mergeAnswers(template []QuestionItem, answers []string) []QuestionItem {
	out := make([]QuestionItem, len(template))
	for i, q := range template {
		out[i] = QuestionItem{
			FieldDesc: q.FieldDesc,
			Content:   q.Content,
			Answer:    answers[i],
		}
	}
	return out
}

func formatConfirmationText(header string, questions []QuestionItem) string {
	var sb strings.Builder
	sb.WriteString(header + "\n\n")
	for _, q := range questions {
		sb.WriteString(q.FieldDesc + ": " + q.Answer + "\n")
	}
	return sb.String()
}
