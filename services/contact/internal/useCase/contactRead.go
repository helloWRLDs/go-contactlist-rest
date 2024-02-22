package usecase

import (
	"context"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
	"helloWRLDs/clean_arch/services/contact/internal/repository"
)

type ContactReadUseCaseImpl struct {
	repository *repository.ContactRepository
}

func NewContactReadUseCase(repository repository.ContactRepository) *ContactReadUseCaseImpl {
	return &ContactReadUseCaseImpl{
		repository: &repository,
	}
}

func (u *ContactUseCaseImpl) List(ctx context.Context) ([]domain.Contact, error) {
	contacts, err := u.repository.GetAll()
	if err != nil {
		return []domain.Contact{}, err
	}
	return contacts, nil
}

func (u *ContactUseCaseImpl) ReadByID(ctx context.Context, id int) (domain.Contact, error) {
	contact, err := u.repository.Get(id)
	if err != nil {
		return domain.Contact{}, err
	}
	return contact, nil
}

func (u *ContactUseCaseImpl) Exists(ctx context.Context, id int) bool {
	return u.repository.Exist(id)
}
