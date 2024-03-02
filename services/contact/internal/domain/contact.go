package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"time"
)

var (
	phoneRegex         = `^(?:\+|[0-9])[0-9]*$`
	ErrEmptyNameFields = errors.New("name fileds cannot be empty")
	ErrNoPhoneNumber   = errors.New("phone number is required")
)

type (
	Contact struct {
		Id         int       `json:"id"`
		FullName   string    `json:"full_name"`
		Phone      string    `json:"phone"`
		CreatedAt  time.Time `json:"created_at"`
		ModifiedAt time.Time `json:"modified_at"`
	}

	ContactDTO struct {
		Id         int       `json:"id"`
		FirstName  string    `json:"first_name"`
		LastName   string    `json:"last_name"`
		Phone      string    `json:"phone"`
		CreatedAt  time.Time `json:"created_at"`
		ModifiedAt time.Time `json:"modified_at"`
	}
)

func NewContact(firstName, lastName, phone string) (*Contact, error) {
	if len(firstName) == 0 || len(lastName) == 0 {
		return &Contact{}, ErrEmptyNameFields
	}
	fullName := fmt.Sprintf("%s %s", firstName, lastName)

	phonePattern := regexp.MustCompile(phoneRegex)
	if !phonePattern.MatchString(phone) {
		return &Contact{}, ErrNoPhoneNumber
	}

	return &Contact{
		Id:         -1,
		FullName:   fullName,
		Phone:      phone,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}, nil
}

func SetContact(id int, fullName, phone string, createdAt, modified time.Time) (*Contact, error) {
	if len(fullName) == 0 {
		return &Contact{}, ErrEmptyNameFields
	}

	phonePattern := regexp.MustCompile(phoneRegex)
	if !phonePattern.MatchString(phone) {
		return &Contact{}, ErrNoPhoneNumber
	}

	return &Contact{
		Id:         id,
		FullName:   fullName,
		Phone:      phone,
		CreatedAt:  createdAt,
		ModifiedAt: modified,
	}, nil
}

func (c *Contact) JSON() []byte {
	json, err := json.Marshal(c)
	if err != nil {
		return nil
	}
	return json
}
