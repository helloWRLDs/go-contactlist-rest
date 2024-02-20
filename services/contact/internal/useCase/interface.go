package usecase

import (
	"context"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
)

type (
	Contact interface {
		Create(ctx context.Context, contact *domain.Contact) (int, error)
		Update(ctx context.Context, id int, contactToUpdate domain.Contact) (int, error)
		Delete(ctx context.Context, id int) error
	}

	ContactReader interface {
		List(ctx context.Context) ([]domain.Contact, error)
		ReadByID(ctx context.Context, id int) (domain.Contact, error)
		Exists(ctx context.Context, id int) bool
	}

	Group interface {
		Create(ctx context.Context, group *domain.Group) (int, error)
		Update(ctx context.Context, id int, groupToUpdate *domain.Group) (int, error)
		Delete(ctx context.Context, id int) error
	}

	GroupRead interface {
		List(ctx context.Context) ([]domain.Group, error)
		ReadByID(ctx context.Context, id int) (domain.Group, error)
		Exists(ctx context.Context, id int) bool
	}

	ContactInGroup interface {
		CreateContactInGroup(ctx context.Context, groupId int, contact *domain.Contact) (int, error)
		AddContactToGroup(ctx context.Context, groupId int, contactId int) error
		DeleteContactFromGroup(ctx context.Context, groupId int, contactId int) error
	}
)
