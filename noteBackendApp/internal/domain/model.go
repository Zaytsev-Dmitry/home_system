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
