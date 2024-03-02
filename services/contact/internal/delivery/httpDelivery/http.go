package httpDelivery

import (
	"database/sql"
	usecase "helloWRLDs/clean_arch/services/contact/internal/useCase"

	"github.com/gin-gonic/gin"
)

type HttpRouter struct {
	Router    *gin.Engine
	ContactUC usecase.ContactUseCase
	GroupUC   usecase.GroupUseCase
}

func NewHttpRouter(db *sql.DB) *HttpRouter {
	return &HttpRouter{
		Router:    gin.New(),
		ContactUC: usecase.NewContactUseCase(db),
		GroupUC:   usecase.NewGroupUseCase(db),
	}
}

func (d *HttpRouter) InitRoutes() {
	d.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	d.contactDelivery()
	d.groupDelivery()
}
