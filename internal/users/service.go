package users

import (
	"context"
)

type Service interface {
	InsertUser(ctx context.Context, user UserSignUp) (User, error)
	UpdateUser(ctx context.Context, user User) (User, error)
	DeleteUser(ctx context.Context, id string) error
	FindUserById(ctx context.Context, id string) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

// DeleteUser implements Service.
func (s *service) DeleteUser(ctx context.Context, id string) error {
	return s.repository.DeleteUser(ctx, id)
}

// FindUserById implements Service.
func (s *service) FindUserById(ctx context.Context, id string) (User, error) {
	user, err := s.repository.FindUserById(ctx, id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// InsertNote implements Service.
func (s *service) InsertUser(ctx context.Context, userReq UserSignUp) (User, error) {
	user, err := HashPassword(AddID(ParseToUser(userReq)))
	if err != nil {
		return User{}, err
	}
	usr, err := s.repository.InsertUser(ctx, user)
	if err != nil {
		return User{}, err
	}
	return usr, nil
}

// UpdateUser implements Service.
func (s *service) UpdateUser(ctx context.Context, user User) (User, error) {
	res, err := s.repository.UpdateUser(ctx, user)
	if err != nil {
		return User{}, err
	}
	return res, nil
}

// FindUserByEmail implements Service.
func (s *service) FindUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := s.repository.FindUserByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
