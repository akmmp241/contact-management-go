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

type AddressControllerImpl struct {
	Service contracts.AddressService
}

func NewAddressControllerImpl(service contracts.AddressService) *AddressControllerImpl {
	return &AddressControllerImpl{Service: service}
}

func (c AddressControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var req dto.AddressDTO
	helper.ReadFromRequestBody(r, &req)

	contactId := params.ByName("contactId")
	id, err := strconv.Atoi(contactId)
	if err != nil {
		panic(exception.NewBadRequest("contactId must be int"))
	}

	addressDTO := c.Service.Create(r.Context(), req, id)

	webResponse := dto.WebResponse{
		Message: "Success create address",
		Data:    addressDTO,
	}

	w.WriteHeader(http.StatusCreated)
	helper.WriteToResponse(w, &webResponse)
}

func (c AddressControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var addressDTO dto.AddressDTO
	helper.ReadFromRequestBody(r, &addressDTO)

	contactParam := params.ByName("contactId")
	contactId, err := strconv.Atoi(contactParam)
	if err != nil {
		panic(exception.NewNotFound("contactId must be int"))
	}

	addressParam := params.ByName("addressId")
	addressId, err := strconv.Atoi(addressParam)
	if err != nil {
		panic(exception.NewBadRequest("addressId must be int"))
	}

	updateAddressResponse := c.Service.Update(r.Context(), addressDTO, contactId, addressId)

	webResponse := dto.WebResponse{
		Message: "Success update address",
		Data:    updateAddressResponse,
	}

	w.WriteHeader(http.StatusOK)
	helper.WriteToResponse(w, &webResponse)
}

func (c AddressControllerImpl) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	contactParam := params.ByName("contactId")
	contactId, err := strconv.Atoi(contactParam)
	if err != nil {
		panic(exception.NewNotFound("contactId must be int"))
	}

	addressParam := params.ByName("addressId")
	addressId, err := strconv.Atoi(addressParam)
	if err != nil {
		panic(exception.NewBadRequest("addressId must be int"))
	}

	getAddressResponse := c.Service.Get(r.Context(), contactId, addressId)

	webResponse := dto.WebResponse{
		Message: "Success get address",
		Data:    getAddressResponse,
	}

	w.WriteHeader(http.StatusOK)
	helper.WriteToResponse(w, &webResponse)
}

func (c AddressControllerImpl) List(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	contactParam := params.ByName("contactId")
	contactId, err := strconv.Atoi(contactParam)
	if err != nil {
		panic(exception.NewNotFound("contactId must be int"))
	}

	listAddressResponse := c.Service.List(r.Context(), contactId)

	webResponse := dto.WebResponse{
		Message: "Success list address",
		Data:    listAddressResponse,
	}

	w.WriteHeader(http.StatusOK)
	helper.WriteToResponse(w, &webResponse)
}

func (c AddressControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	contactParam := params.ByName("contactId")
	contactId, err := strconv.Atoi(contactParam)
	if err != nil {
		panic(exception.NewNotFound("contactId must be int"))
	}

	addressParam := params.ByName("addressId")
	addressId, err := strconv.Atoi(addressParam)
	if err != nil {
		panic(exception.NewBadRequest("addressId must be int"))
	}

	c.Service.Delete(r.Context(), contactId, addressId)

	webResponse := dto.WebResponse{
		Message: "Success delete address",
		Data:    true,
	}

	w.WriteHeader(http.StatusOK)
	helper.WriteToResponse(w, &webResponse)
}
