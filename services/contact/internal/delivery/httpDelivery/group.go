package httpDelivery

func (d *HttpRouter) groupDelivery() {
	groupRouter := d.Router.Group("/groups")
	{
		groupRouter.GET("/", d.GroupUC.GetAllGroups)

		groupRouter.GET("/:id", d.GroupUC.GetGroup)

		groupRouter.POST("/", d.GroupUC.InsertGroup)

		groupRouter.DELETE("/:id", d.GroupUC.DeleteGroup)

		// groupRouter.POST("/:id", d.GroupUC.UpdateGroup)
	}
}
