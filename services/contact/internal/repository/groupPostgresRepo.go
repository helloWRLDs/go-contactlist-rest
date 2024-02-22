package repository

import (
	"database/sql"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
)

type GroupRepositoryImpl struct {
	DB *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{
		DB: db,
	}
}

func (r *GroupRepositoryImpl) Insert(group *domain.Group) (int, error) {
	var id int
	stmt := "INSERT INTO groups(names) VALUES ($1) RETURNING id"
	err := r.DB.QueryRow(stmt, group.Name).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *GroupRepositoryImpl) Get(id int) (domain.Group, error) {
	var group domain.Group
	stmt := "SELECT * FROM groups WHERE id=$1"
	err := r.DB.QueryRow(stmt, id).Scan(&group.Id, &group.Name)
	if err != nil {
		return domain.Group{}, err
	}
	return group, nil
}

func (r *GroupRepositoryImpl) GetAll() ([]domain.Group, error) {
	var groups []domain.Group
	stmt := "SELECT * FROM groups"
	rows, err := r.DB.Query(stmt)
	if err != nil {
		return []domain.Group{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var group domain.Group
		err = rows.Scan(&group.Id, &group.Name)
		if err != nil {
			return []domain.Group{}, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (r *GroupRepositoryImpl) Delete(id int) error {
	stmt := "DELETE FROM groups WHERE id=$1"
	_, err := r.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *GroupRepositoryImpl) Exist(id int) bool {
	stmt := "SELECT EXISTS (SELECT * FROM groups WHERE id=$1)"
	var exists bool
	err := r.DB.QueryRow(stmt, id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
