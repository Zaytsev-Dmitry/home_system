package echo

type Result struct {
	MsgId   int
	ChatId  int64
	Answers []CollectItem
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
