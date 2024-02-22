package usecase

import (
	"context"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
	"helloWRLDs/clean_arch/services/contact/internal/repository"
)

type ContactUseCaseImpl struct {
	repository repository.ContactRepository
}

func NewContactUseCase(repository *repository.ContactRepositoryImpl) *ContactUseCaseImpl {
	return &ContactUseCaseImpl{
		repository: repository,
	}
}

func (u *ContactUseCaseImpl) Create(ctx context.Context, contact *domain.Contact) (int, error) {
	if err := contact.Validate(); err != nil {
		return 0, err
	}
	id, err := u.repository.Insert(contact)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *ContactUseCaseImpl) Update(ctx context.Context, id int, contactToUpdate domain.Contact) (int, error) {
	return 0, nil
}

func (u *ContactUseCaseImpl) Delete(ctx context.Context, id int) error {
	return nil
}
