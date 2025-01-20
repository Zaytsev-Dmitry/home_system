package noteDomain

type NoteEntity struct {
	Id   string
	Name string
	Link string
}

type CreateNoteRequest struct {
	Name string `json:"name"`
	Link string `json:"link,optional"`
}

type NoteResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Link string `json:"link,optional"`
}

type ResponseError struct {
	Timestamp    string `json:"timestamp"`
	Status       int    `json:"status"`
	BusinessCode string `json:"businessCode,omitempty"`
	Error        string `json:"error"`
	Path         string `json:"path"`
}
