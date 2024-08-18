package impl

import (
	"contact-management-restful/exception"
	"contact-management-restful/models/dto"
	"contact-management-restful/services/contracts"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ContactControllerImpl struct {
	ContactService contracts.ContactService
}

func NewContactControllerImpl(contactService contracts.ContactService) *ContactControllerImpl {
	return &ContactControllerImpl{ContactService: contactService}
}

func (c ContactControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var createContactRequest dto.CreateContactRequest
	err := json.NewDecoder(r.Body).Decode(&createContactRequest)
	if err != nil {
		panic(exception.NewBadRequest("invalid request body. Content-Type must be application/json"))
	}

	createContactResponse := c.ContactService.Create(r.Context(), createContactRequest)

	webResponse := dto.WebResponse{
		Message: "Success Create contact",
		Data:    createContactResponse,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(webResponse); err != nil {
		panic(err)
	}
}

func (c ContactControllerImpl) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	searchContactResponse := c.ContactService.Search(r.Context())

	webResponse := dto.WebResponse{
		Message: "Success get contact list",
		Data:    searchContactResponse,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(webResponse); err != nil {
		panic(err)
	}
}

func (c ContactControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var updateContactRequest dto.UpdateContactRequest
	err := json.NewDecoder(r.Body).Decode(&updateContactRequest)
	if err != nil {
		panic(exception.NewBadRequest("invalid request body. Content-Type must be application/json"))
	}

	idFromParam := params.ByName("id")
	id, err := strconv.Atoi(idFromParam)
	if err != nil {
		panic(exception.NewBadRequest("invalid path parameter"))
	}

	createContactResponse := c.ContactService.Update(r.Context(), updateContactRequest, id)

	webResponse := dto.WebResponse{
		Message: "Success Update contact",
		Data:    createContactResponse,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(webResponse); err != nil {
		panic(err)
	}
}

func (c ContactControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idFromParam := params.ByName("id")
	id, err := strconv.Atoi(idFromParam)
	if err != nil {
		panic(exception.NewNotFound("invalid path parameter"))
	}

	getContactsResponse := c.ContactService.Get(r.Context(), id)

	webResponse := dto.WebResponse{
		Message: "Success Get contact",
		Data:    getContactsResponse,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(webResponse); err != nil {
		panic(err)
	}
}

func (c ContactControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idFromParam := params.ByName("id")
	id, err := strconv.Atoi(idFromParam)
	if err != nil {
		panic(exception.NewNotFound("invalid path parameter"))
	}

	c.ContactService.Delete(r.Context(), id)

	webResponse := dto.WebResponse{
		Message: "Success Delete contact",
		Data:    true,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(webResponse); err != nil {
		panic(err)
	}
}
