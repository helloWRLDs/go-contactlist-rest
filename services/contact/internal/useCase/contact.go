package usecase

import (
	"context"
	"database/sql"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
	"helloWRLDs/clean_arch/services/contact/internal/repository"
)

type ContactUseCase struct {
	repo *repository.ContactRepository
}

func NewContactUseCase(db *sql.DB) *ContactUseCase {
	return &ContactUseCase{
		repo: repository.NewContactRepository(db),
	}
}

func (u *ContactUseCase) Create(ctx context.Context, contact *domain.Contact) (int, error) {
	return 0, nil
}

func (u *ContactUseCase) Update(ctx context.Context, id int, contactToUpdate domain.Contact) (int, error) {
	return 0, nil
}

func (u *ContactUseCase) Delete(ctx context.Context, id int) error {
	return nil
}
