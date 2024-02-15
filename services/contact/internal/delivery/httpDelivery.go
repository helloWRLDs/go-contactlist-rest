package delivery

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"helloWRLDs/clean_arch/services/contact/internal/domain"
	usecase "helloWRLDs/clean_arch/services/contact/internal/useCase"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type (
	HTTPDelivery struct {
		UseCase usecase.UseCaseInterface
		Logger  *slog.Logger
	}

	HTTPDeliveryInterface interface {
		GetAllContactsController(w http.ResponseWriter, r *http.Request)
		GetContactController(w http.ResponseWriter, r *http.Request)
		InsertContactController(w http.ResponseWriter, r *http.Request)
		DeleteContactController(w http.ResponseWriter, r *http.Request)
		// UpdateContactController(w http.ResponseWriter, r *http.Request)
	}
)

func NewDelivery(db *sql.DB, logger *slog.Logger) *HTTPDelivery {
	return &HTTPDelivery{usecase.NewUseCase(db), logger}
}

func (d *HTTPDelivery) GetAllContactsController(w http.ResponseWriter, r *http.Request) {
	var contacts []domain.Contact
	contacts, err := d.UseCase.RetrieveAllContactsUsecase()
	if err != nil {
		d.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	contactsJSON, err := json.Marshal(contacts)
	if err != nil {
		d.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Write(contactsJSON)
}

func (d *HTTPDelivery) GetContactController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("id"))
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Wrong query(no id in query)"))
		return
	}
	contact, err := d.UseCase.RetrieveContactUsecase(id)
	if err != nil {
		d.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("db error"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(contact.JSON()))
}

func (d *HTTPDelivery) InsertContactController(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		d.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var contact domain.Contact
	err = json.Unmarshal(body, &contact)
	if err != nil {
		d.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := d.UseCase.InsertContactUsecase(contact)
	if err != nil {
		d.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Contact inserted with id=%d", id)))
}

func (d *HTTPDelivery) DeleteContactController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Wrong query(no id in query)"))
		return
	}
	err = d.UseCase.DeleteContactUsecase(id)
	if err != nil {
		d.Logger.Error(err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Wrong query(no id in query)"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted contact with id=%d", id)))
}

// func (d *HTTPDelivery) UpdateContactController(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("id"))
// 	if err != nil {
// 		w.WriteHeader(500)
// 		w.Write([]byte("Wrong query(no id in query)"))
// 		return
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(500)
// 		return
// 	}
// 	var contactToUpdate domain.Contact
// 	err = json.Unmarshal(body, &contactToUpdate)
// 	if err != nil {
// 		w.WriteHeader(500)
// 		w.Write([]byte("Body parsing error"))
// 		return
// 	}
// 	d.UseCase.UpdateContactUsecase(id, contactToUpdate)
// }
