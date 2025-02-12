package echo

import "telegramCLient/internal/storage"

type Result struct {
	Answers  []CollectItem
	Messages []storage.Message
}

type CollectItem struct {
	FieldId   string
	FieldName string
	Content   string
	Answer    string
}
