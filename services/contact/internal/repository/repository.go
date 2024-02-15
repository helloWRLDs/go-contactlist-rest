package repository

import (
	"database/sql"
	"fmt"
	"helloWRLDs/clean_arch/services/contact/internal/domain"

	_ "github.com/lib/pq"
)

type (
	Repository struct {
		DB *sql.DB
	}

	RepositoryInterface interface {
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
)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// contact basic CRUD operations
func (r *Repository) GetContact(id int) (domain.Contact, error) {
	var c domain.Contact
	stmt := "SELECT * FROM contacts WHERE id=$1"
	err := r.DB.QueryRow(stmt, id).Scan(
		&c.Id,
		&c.FirstName,
		&c.LastName,
		&c.MiddleName,
		&c.Phone,
		&c.GroupId,
	)
	if err != nil {
		return domain.Contact{}, err
	}
	return c, nil
}
func (r *Repository) GetAllContacts() ([]domain.Contact, error) {
	stmt := "SELECT * FROM contacts"
	rows, err := r.DB.Query(stmt)
	if err != nil {
		return []domain.Contact{}, err
	}
	defer rows.Close()

	var contacts []domain.Contact
	for rows.Next() {
		var c domain.Contact
		err = rows.Scan(&c.Id, &c.FirstName, &c.LastName, &c.MiddleName, &c.Phone, &c.GroupId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}
func (r *Repository) InsertContact(c domain.Contact) (int, error) {
	stmt := "INSERT INTO contacts (first_name, last_name, middle_name, phone, group_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	err := r.DB.QueryRow(stmt, c.FirstName, c.LastName, c.MiddleName, c.Phone, c.GroupId).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (r *Repository) DeleteContact(id int) error {
	stmt := "DELETE FROM contacts WHERE id=$1"
	_, err := r.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

// group basic CRUD operations
func (r *Repository) GetGroup(id int) (domain.Group, error) {
	var group domain.Group
	stmt := "SELECT * FROM groups WHERE id=$1"
	err := r.DB.QueryRow(stmt, id).Scan(&group.Id, &group.Name)
	if err != nil {
		return domain.Group{}, err
	}
	return group, nil
}
func (r *Repository) GetAllGroups() ([]domain.Group, error) {
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
func (r *Repository) InsertGroup(g domain.Group) (int, error) {
	var id int
	stmt := "INSERT INTO groups(names) VALUES ($1) RETURNING id"
	err := r.DB.QueryRow(stmt, g.Name).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (r *Repository) DeleteGroup(id int) error {
	stmt := "DELETE FROM groups WHERE id=$1"
	_, err := r.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

// specific CRUD operations
func (r *Repository) GetContactsByGroup(group_id int) ([]domain.Contact, error) {
	var contacts []domain.Contact
	stmt := "SELECT * FROM contacts WHERE group_id=$1"
	rows, err := r.DB.Query(stmt, group_id)
	if err != nil {
		return []domain.Contact{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var contact domain.Contact
		err = rows.Scan(
			&contact.Id,
			&contact.FirstName,
			&contact.LastName,
			&contact.MiddleName,
			&contact.Phone,
			&contact.GroupId,
		)
		if err != nil {
			return []domain.Contact{}, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}
func (r *Repository) IsContactExist(id int) bool {
	stmt := "SELECT EXISTS (SELECT * FROM contacts WHERE id=$1)"
	var exists bool
	err := r.DB.QueryRow(stmt, id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
