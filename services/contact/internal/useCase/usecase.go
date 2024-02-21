package usecase

// import (
// 	"database/sql"
// 	"fmt"
// 	"helloWRLDs/clean_arch/services/contact/internal/domain"
// 	"helloWRLDs/clean_arch/services/contact/internal/repository"
// 	"reflect"
// )

// type (
// 	UseCase struct {
// 		repo repository.RepositoryInterface
// 	}

// 	UseCaseInterface interface {
// 		RetrieveAllContactsUsecase() ([]domain.Contact, error)
// 		RetrieveContactUsecase(id int) (domain.Contact, error)
// 		InsertContactUsecase(contact domain.Contact) (int, error)
// 		DeleteContactUsecase(id int) error
// 	}
// )

// func NewUseCase(db *sql.DB) *UseCase {
// 	return &UseCase{repo: repository.NewRepository(db)}
// }

// func (u *UseCase) RetrieveAllContactsUsecase() ([]domain.Contact, error) {
// 	var contacts []domain.Contact
// 	contacts, err := u.repo.GetAllContacts()
// 	if err != nil {
// 		return []domain.Contact{}, err
// 	}
// 	return contacts, nil
// }

// func (u *UseCase) RetrieveContactUsecase(id int) (domain.Contact, error) {
// 	contact, err := u.repo.GetContact(id)
// 	if err != nil {
// 		return domain.Contact{}, err
// 	}
// 	return contact, nil
// }

// func (u *UseCase) InsertContactUsecase(contact domain.Contact) (int, error) {
// 	t := reflect.TypeOf(contact)
// 	// Iterating through fields in contact object
// 	for i := 0; i < t.NumField(); i++ {
// 		fieldName := t.Field(i).Name
// 		fieldValue := reflect.ValueOf(contact).Field(i).Interface()
// 		// Set the groupId value to 0 (none) if there is no specified group
// 		if fieldName == "GroupId" && fieldValue == 0 {
// 			contact.GroupId = 1
// 		}
// 		// Check if name fields have contain null values
// 		if (fieldName == "MiddleName" || fieldName == "FirstName" || fieldName == "LastName") && fieldValue == "" {
// 			return 0, fmt.Errorf("name fields cannot be a null")
// 		}
// 	}
// 	id, err := u.repo.InsertContact(contact)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }

// func (u *UseCase) DeleteContactUsecase(id int) error {
// 	if u.repo.IsContactExist(id) {
// 		err := u.repo.DeleteContact(id)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// 	return fmt.Errorf("user doesn't exist")
// }
