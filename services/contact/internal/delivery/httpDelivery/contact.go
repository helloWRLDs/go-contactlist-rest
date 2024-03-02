package httpDelivery

type ContactDelivery interface {
}

type ContactDeliveryImpl struct {
}

func (d *HttpRouter) contactDelivery() {
	contactRouter := d.Router.Group("/contacts")
	{
		contactRouter.GET("/", d.ContactUC.GetAllContacts)

		contactRouter.POST("/", d.ContactUC.InsertContact)

		contactRouter.GET("/:id", d.ContactUC.GetContact)

		contactRouter.PUT("/:id", d.ContactUC.UpdateContact)

		contactRouter.DELETE("/:id", d.ContactUC.DeleteContact)
	}
}
