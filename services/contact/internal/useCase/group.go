package usecase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
	"helloWRLDs/clean_arch/services/contact/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupUseCaseImpl struct {
	repo repository.GroupRepository
}

func NewGroupUseCase(db *sql.DB) *GroupUseCaseImpl {
	return &GroupUseCaseImpl{
		repo: repository.NewGroupRepository(db),
	}
}

func (u *GroupUseCaseImpl) GetAllGroups(ctx *gin.Context) {
	groups, err := u.repo.GetAll()
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	groupsJson, err := json.Marshal(groups)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	ctx.Writer.WriteHeader(200)
	ctx.Writer.Write(groupsJson)
}

func (u *GroupUseCaseImpl) InsertGroup(ctx *gin.Context) {
	var group domain.Group
	if err := ctx.BindJSON(&group); err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	if !group.Validate() {
		ctx.Writer.WriteHeader(500)
		fmt.Println()
		return
	}
	newGroup, err := domain.NewGroup(group.Name, group.Description)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	id, err := u.repo.Insert(newGroup)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	ctx.Writer.WriteHeader(201)
	ctx.Writer.Write([]byte(fmt.Sprintf("inserted contact with id=%d", id)))
}

func (u *GroupUseCaseImpl) GetGroup(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 1 {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	group, err := u.repo.Get(id)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		fmt.Println(err)
		return
	}
	ctx.Writer.WriteHeader(200)
	ctx.Writer.Write([]byte(group.JSON()))
}

func (u *GroupUseCaseImpl) DeleteGroup(ctx *gin.Context) {
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
	ctx.Writer.Write([]byte(fmt.Sprintf("deleted group with id=%d", id)))
}
