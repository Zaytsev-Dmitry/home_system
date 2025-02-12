package echo

import "telegramCLient/internal/storage"

type Result struct {
	MsgId         int
	ChatId        int64
	UserFirstName string
	UserLastname  string
	UserTGName    string
	Answers       []CollectItem
	Messages      []storage.Message
}
type CollectItem struct {
	FieldId   string
	FieldName string
	Content   string
}

type dataCollect struct {
	State   State
	answers []CollectItem
}
