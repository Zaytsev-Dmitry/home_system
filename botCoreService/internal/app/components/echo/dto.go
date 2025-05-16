package echo

type TextMeta struct {
	ConfirmText string
	StartText   string
}

type QuestionItem struct {
	FieldDesc string
	Content   string
	Answer    string
}

type EchoResult struct {
	ChatId   int64
	Question []QuestionItem
}
