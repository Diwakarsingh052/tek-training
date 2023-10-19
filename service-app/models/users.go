package models

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) (*Service, error) { //db *sql.DB
	if db == nil {
		return nil, errors.New("please provide a valid connection")
	}
	s := &Service{db: db}
	return s, nil
}

func (s *Service) CreateUser(ctx context.Context, nu NewUser, now time.Time) (User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generating password hash %w", err)
	}
	u := User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: string(hashedPass),
		DateCreated:  now,
		DateUpdated:  now,
	}

	err = s.db.Create(&u).Error
	if err != nil {
		return User{}, err
	}

	return u, nil

}
