package echo

type Result struct {
	ChatId        int64
	UserFirstName string
	UserLastname  string
	Username      string
	Question      []CollectItem
	MessagesIds   []int
}

type CollectItem struct {
	FieldId   string
	FieldName string
	Content   string
	Answer    string
}
