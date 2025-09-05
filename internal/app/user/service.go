package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(username, email, password string) (*User, error)
	Login(email, password string) (*User, error)
	ChangePassword(userID uint, oldPass, newPass string) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Register(username, email, password string) (*User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	u := &User{Username: username, Email: email, Password: string(hashed)}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) Login(email, password string) (*User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return &u, nil
}

func (s *service) ChangePassword(userID uint, oldPass, newPass string) error {
	u, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(oldPass)) != nil {
		return errors.New("old password incorrect")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(newPass), 12)
	return s.repo.UpdatePassword(userID, string(hash))
}
