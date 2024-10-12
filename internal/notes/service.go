package notes

import (
	"context"
)

type Service interface {
	InsertNote(ctx context.Context, note Note) (Note, error)
	UpdateNote(ctx context.Context, note Note) (Note, error)
	DeleteNote(ctx context.Context, id string) error
	FindNoteById(ctx context.Context, id string) (Note, error)
	GetNotes(ctx context.Context, user_id string) ([]Note, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

// DeleteNote implements Adapter.
func (s *service) DeleteNote(ctx context.Context, id string) error {
	return s.repository.DeleteNote(ctx, id)
}

// FindNoteById implements Adapter.
func (s *service) FindNoteById(ctx context.Context, id string) (Note, error) {
	note, err := s.repository.FindNoteById(ctx, id)
	if err != nil {
		return Note{}, err
	}
	return note, nil
}

// GetNotes implements Adapter.
func (s *service) GetNotes(ctx context.Context, user_id string) ([]Note, error) {
	notes, err := s.repository.GetNotes(ctx, user_id)
	if err != nil {
		return []Note{}, err
	}
	return notes, nil
}

// InsertNote implements Adapter.
func (s *service) InsertNote(ctx context.Context, note Note) (Note, error) {
	newNote, err := s.repository.InsertNote(ctx, AddID(note))
	if err != nil {
		return Note{}, err
	}
	return newNote, nil
}

// UpdateNote implements Adapter.
func (s *service) UpdateNote(ctx context.Context, note Note) (Note, error) {
	res, err := s.repository.UpdateNote(ctx, note)
	if err != nil {
		return Note{}, err
	}
	return res, nil
}
