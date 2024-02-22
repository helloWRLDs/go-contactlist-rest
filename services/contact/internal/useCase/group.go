package usecase

import (
	"context"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
	"helloWRLDs/clean_arch/services/contact/internal/repository"
)

type GroupUseCaseImpl struct {
	repository repository.GroupRepository
}

func NewGroupRepository(repository *repository.GroupRepositoryImpl) *GroupUseCaseImpl {
	return &GroupUseCaseImpl{
		repository: repository,
	}
}

func (r *GroupUseCaseImpl) Create(ctx context.Context, group *domain.Group) (int, error) {
	if err := group.Validate(); err != nil {
		return 0, err
	}
	r.
	return
}

func (r *GroupUseCaseImpl) Update(ctx context.Context, id int, groupToUpdate *domain.Group) (int, error) {
	return
}

func (r *GroupUseCaseImpl) Delete(ctx context.Context, id int) error {
	return
}
