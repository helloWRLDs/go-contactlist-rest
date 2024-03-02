package domain

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrEmptyName        = errors.New("group name cannot be empty")
	ErrEmptyDescription = errors.New("group description cannot be empty")
)

type Group struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

func NewGroup(name, description string) (*Group, error) {
	if len(name) == 0 {
		return &Group{}, ErrEmptyName
	}
	if len(description) == 0 {
		return &Group{}, ErrEmptyDescription
	}
	return &Group{
		Id:          -1,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}, nil
}

func SetGroup(id int, name, description string, createdAt, modifiedAt time.Time) (*Group, error) {
	if len(name) == 0 {
		return &Group{}, ErrEmptyName
	}
	if len(description) == 0 {
		return &Group{}, ErrEmptyDescription
	}
	return &Group{
		Id:          id,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
		ModifiedAt:  modifiedAt,
	}, nil
}

func (g *Group) JSON() []byte {
	json, err := json.Marshal(g)
	if err != nil {
		return nil
	}
	return json
}

func (g *Group) Validate() bool {
	if len(g.Name) == 0 || len(g.Description) == 0 {
		return false
	}
	return true
}
