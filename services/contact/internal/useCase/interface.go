package usecase

import "github.com/gin-gonic/gin"

type (
	ContactUseCase interface {
		GetAllContacts(ctx *gin.Context)
		InsertContact(ctx *gin.Context)
		GetContact(ctx *gin.Context)
		DeleteContact(ctx *gin.Context)
		UpdateContact(ctx *gin.Context)
	}

	GroupUseCase interface {
		GetAllGroups(ctx *gin.Context)
		InsertGroup(ctx *gin.Context)
		GetGroup(ctx *gin.Context)
		DeleteGroup(ctx *gin.Context)
	}
)
