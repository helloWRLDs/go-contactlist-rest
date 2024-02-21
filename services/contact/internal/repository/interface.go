package repository

import "helloWRLDs/clean_arch/services/contact/internal/domain"

type (
	ContactInterface interface {
		Insert(contact *domain.Contact) (int, error)
		Get(id int) (domain.Contact, error)
		GetAll() ([]domain.Contact, error)
		Delete(id int) error
		Exist(id int) bool
	}

	GroupInterface interface {
		Insert(group *domain.Group) (int, error)
		Get(id int) (domain.Group, error)
		GetAll() ([]domain.Group, error)
		Delete(id int) error
		Exist(id int) bool
	}
)
