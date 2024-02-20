package repository

import "helloWRLDs/clean_arch/services/contact/internal/domain"

type RepositoryInterface interface {
	IsContactExist(id int) bool
	GetContact(id int) (domain.Contact, error)
	GetAllContacts() ([]domain.Contact, error)
	InsertContact(c domain.Contact) (int, error)
	DeleteContact(id int) error
	GetGroup(id int) (domain.Group, error)
	GetAllGroups() ([]domain.Group, error)
	InsertGroup(c domain.Group) (int, error)
	DeleteGroup(id int) error
	GetContactsByGroup(group_id int) ([]domain.Contact, error)
}

type (
	Contact interface {
		Insert()
	}
)
