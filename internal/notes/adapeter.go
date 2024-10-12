package notes

import (
	"github.com/google/uuid"
)

func AddID(note Note) Note {
	note.ID = uuid.New().String()
	return note
}
