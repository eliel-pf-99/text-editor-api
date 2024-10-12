package notes

import "context"

type Handler interface {
	InsertNote(ctx context.Context, note Note) (Note, error)
	UpdateNote(ctx context.Context, note Note) (Note, error)
	DeleteNote(ctx context.Context, id string) error
	FindNoteById(ctx context.Context, id string) (Note, error)
	GetNotes(ctx context.Context, user_id string) ([]Note, error)
}

type handler struct {
	service Service
}

// DeleteNote implements Handler.
func (h *handler) DeleteNote(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindNoteById implements Handler.
func (h *handler) FindNoteById(ctx context.Context, id string) (Note, error) {
	panic("unimplemented")
}

// GetNotes implements Handler.
func (h *handler) GetNotes(ctx context.Context, user_id string) ([]Note, error) {
	panic("unimplemented")
}

// InsertNote implements Handler.
func (h *handler) InsertNote(ctx context.Context, note Note) (Note, error) {
	h.service.InsertNote(ctx, note)
	return Note{}, nil
}

// UpdateNote implements Handler.
func (h *handler) UpdateNote(ctx context.Context, note Note) (Note, error) {
	panic("unimplemented")
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}
