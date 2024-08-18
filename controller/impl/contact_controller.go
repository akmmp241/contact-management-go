package impl

import (
	"contact-management-restful/exception"
	"contact-management-restful/helper"
	"contact-management-restful/models/dto"
	"contact-management-restful/services/contracts"
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
	helper.ReadFromRequestBody(r, &createContactRequest)

	createContactResponse := c.ContactService.Create(r.Context(), createContactRequest)

	webResponse := dto.WebResponse{
		Message: "Success Create contact",
		Data:    createContactResponse,
	}

	w.WriteHeader(http.StatusCreated)
	helper.WriteToResponse(w, &webResponse)
}

func (c ContactControllerImpl) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	searchContactResponse := c.ContactService.Search(r.Context())

	webResponse := dto.WebResponse{
		Message: "Success get contact list",
		Data:    searchContactResponse,
	}

	w.WriteHeader(http.StatusOK)
	helper.WriteToResponse(w, &webResponse)
}

func (c ContactControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var updateContactRequest dto.UpdateContactRequest
	helper.ReadFromRequestBody(r, &updateContactRequest)

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
	helper.WriteToResponse(w, &webResponse)
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
	helper.WriteToResponse(w, &webResponse)
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
	helper.WriteToResponse(w, &webResponse)
}
