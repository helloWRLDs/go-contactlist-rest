package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	def := alice.New(contentTypeJSON, secureHeaders, app.logRequest)

	router.Handler(http.MethodPost, "/contacts/", def.ThenFunc(app.delivery.InsertContactController))
	router.Handler(http.MethodGet, "/contacts", def.ThenFunc(app.delivery.GetAllContactsController))
	router.Handler(http.MethodGet, "/contacts/:id", def.ThenFunc(app.delivery.GetContactController))
	router.Handler(http.MethodDelete, "/contacts/:id", def.ThenFunc(app.delivery.DeleteContactController))
	// router.HandlerFunc(http.MethodPut, "/contact/:id", app.delivery.UpdateContactController)

	return router
}
