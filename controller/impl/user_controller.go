package impl

import (
	"contact-management-restful/helper"
	"contact-management-restful/models/domains"
	"contact-management-restful/models/dto"
	servicesContracts "contact-management-restful/services/contracts"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserControllerImpl struct {
	UserService servicesContracts.UserService
}

func NewUserController(userService servicesContracts.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var registerRequest dto.RegisterRequest
	helper.ReadFromRequestBody(r, &registerRequest)

	registerResponse := c.UserService.Register(r.Context(), registerRequest)

	webResponse := &dto.WebResponse{
		Message: "Success Register User",
		Data:    registerResponse,
	}

	w.WriteHeader(http.StatusCreated)
	helper.WriteToResponse(w, &webResponse)
}

func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginRequest dto.LoginRequest
	helper.ReadFromRequestBody(r, &loginRequest)

	loginResponse := c.UserService.Login(r.Context(), loginRequest)

	webResponse := &dto.WebResponse{
		Message: "Success Login User",
		Data:    loginResponse,
	}

	w.WriteHeader(http.StatusOK)
	helper.WriteToResponse(w, &webResponse)
}

func (c *UserControllerImpl) Current(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := r.Context().Value("user").(domains.User)

	webResponse := &dto.WebResponse{
		Message: "Success get user data",
		Data: &dto.CurrentResponse{
			Name:     user.Name,
			Username: user.Username,
			Token:    user.Token,
		},
	}

	w.WriteHeader(http.StatusOK)
	helper.WriteToResponse(w, &webResponse)
}

func (c *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var updateRequest dto.UpdateUserRequest
	helper.ReadFromRequestBody(r, &updateRequest)

	updateResponse := c.UserService.Update(r.Context(), updateRequest)

	webResponse := &dto.WebResponse{
		Message: "Success Update User",
		Data:    updateResponse,
	}

	w.WriteHeader(http.StatusOK)
	helper.WriteToResponse(w, &webResponse)
}

func (c *UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c.UserService.Logout(r.Context())

	webResponse := &dto.WebResponse{
		Message: "Success Logout User",
		Data:    true,
	}

	w.WriteHeader(http.StatusOK)
	helper.WriteToResponse(w, &webResponse)
}
