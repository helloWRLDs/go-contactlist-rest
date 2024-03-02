package usecase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
	"helloWRLDs/clean_arch/services/contact/internal/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ContactUseCaseImpl struct {
	repo repository.ContactRepository
}

func NewContactUseCase(db *sql.DB) *ContactUseCaseImpl {
	return &ContactUseCaseImpl{
		repo: repository.NewContactRepository(db),
	}
}

func (u *ContactUseCaseImpl) GetAllContacts(ctx *gin.Context) {
	contacts, err := u.repo.GetAll()
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}

	contactsJson, err := json.Marshal(contacts)
	if err != nil {
		fmt.Println(err)
		ctx.Writer.WriteHeader(500)
		return
	}

	ctx.Writer.WriteHeader(200)
	ctx.Writer.Write(contactsJson)
}

func (u *ContactUseCaseImpl) InsertContact(ctx *gin.Context) {
	var dto domain.ContactDTO
	if err := ctx.BindJSON(&dto); err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}

	contact, err := domain.NewContact(dto.FirstName, dto.LastName, dto.Phone)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	id, err := u.repo.Insert(contact)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	ctx.Writer.WriteHeader(201)
	ctx.Writer.Write([]byte(fmt.Sprintf("inserted contact with id=%d", id)))
}

func (u *ContactUseCaseImpl) GetContact(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 1 {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	contact, err := u.repo.Get(id)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	ctx.Writer.WriteHeader(200)
	ctx.Writer.Write([]byte(contact.JSON()))
}

func (u *ContactUseCaseImpl) DeleteContact(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 1 {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	if !u.repo.Exist(id) {
		ctx.Writer.WriteHeader(404)
		fmt.Println(err)
		return
	}
	if err = u.repo.Delete(id); err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	ctx.Writer.WriteHeader(200)
	ctx.Writer.Write([]byte(fmt.Sprintf("deleted user with id=%d", id)))
}

func (u *ContactUseCaseImpl) UpdateContact(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 1 {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	var updatedContact domain.Contact
	if err = ctx.BindJSON(&updatedContact); err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	updatedContact.ModifiedAt = time.Now()
	resultContact, err := u.repo.Update(id, &updatedContact)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}

	ctx.Writer.WriteHeader(200)
	ctx.Writer.Write(resultContact.JSON())
}
