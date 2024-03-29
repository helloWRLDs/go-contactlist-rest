package repository

import (
	"database/sql"
	"fmt"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
)

type ContactRepositoryImpl struct {
	DB *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepositoryImpl {
	return &ContactRepositoryImpl{
		DB: db,
	}
}

func (r *ContactRepositoryImpl) Update(id int, updatedContact *domain.Contact) (domain.Contact, error) {
	stmt := "UPDATE contacts SET full_name=$1, phone=$2, modified_at=$3 WHERE id=$4 RETURNING *"
	var c domain.Contact
	err := r.DB.QueryRow(stmt,
		updatedContact.FullName,
		updatedContact.Phone,
		updatedContact.ModifiedAt,
		id).Scan(&c.Id, &c.FullName, &c.Phone, &c.CreatedAt, &c.ModifiedAt)
	if err != nil {
		return domain.Contact{}, err
	}
	return c, nil
}

func (r *ContactRepositoryImpl) Insert(contact *domain.Contact) (int, error) {
	stmt := "INSERT INTO contacts (full_name, phone, created_at, modified_at) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := r.DB.QueryRow(stmt, contact.FullName, contact.Phone, contact.CreatedAt, contact.ModifiedAt).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *ContactRepositoryImpl) Get(id int) (domain.Contact, error) {
	var c domain.Contact
	stmt := "SELECT * FROM contacts WHERE id=$1"
	err := r.DB.QueryRow(stmt, id).Scan(
		&c.Id,
		&c.FullName,
		&c.Phone,
		&c.CreatedAt,
		&c.ModifiedAt,
	)
	if err != nil {
		return domain.Contact{}, err
	}
	return c, nil
}

func (r *ContactRepositoryImpl) GetAll() ([]domain.Contact, error) {
	stmt := "SELECT * FROM contacts"
	rows, err := r.DB.Query(stmt)
	if err != nil {
		return []domain.Contact{}, err
	}
	defer rows.Close()

	var contacts []domain.Contact
	for rows.Next() {
		var c domain.Contact
		err = rows.Scan(&c.Id, &c.FullName, &c.Phone, &c.CreatedAt, &c.ModifiedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}

func (r *ContactRepositoryImpl) Delete(id int) error {
	stmt := "DELETE FROM contacts WHERE id=$1"
	_, err := r.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepositoryImpl) Exist(id int) bool {
	stmt := "SELECT EXISTS (SELECT * FROM contacts WHERE id=$1)"
	var exists bool
	err := r.DB.QueryRow(stmt, id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
