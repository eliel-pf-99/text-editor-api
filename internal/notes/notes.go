package notes

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	User_id string `json:"user_id"`
}

type NoteCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteReq struct {
	NoteID string `json:"note_id"`
}
