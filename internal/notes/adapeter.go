package notes

import (
	"github.com/google/uuid"
)

func AddID(note Note) Note {
	note.ID = uuid.New().String()
	return note
}

func AddUserId(noteCreate NoteCreate, userId string) Note {
	var note Note
	note.Title = noteCreate.Title
	note.Content = noteCreate.Content
	note.User_id = userId
	return note
}
