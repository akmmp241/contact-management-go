package impl

import (
	"contact-management-restful/exception"
	"contact-management-restful/models/domains"
	"contact-management-restful/models/dto"
	servicesContracts "contact-management-restful/services/contracts"
	"encoding/json"
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
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		panic(exception.NewBadRequest("invalid request body"))
	}

	registerResponse := c.UserService.Register(r.Context(), registerRequest)

	webResponse := &dto.WebResponse{
		Message: "Success Register User",
		Data:    registerResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		panic(exception.NewBadRequest("invalid request body"))
	}

	loginResponse := c.UserService.Login(r.Context(), loginRequest)

	webResponse := &dto.WebResponse{
		Message: "Success Login User",
		Data:    loginResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(webResponse)
	if err != nil {
		panic(err)
	}
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (c *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var updateRequest dto.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		panic(exception.NewBadRequest("invalid request body"))
	}

	updateResponse := c.UserService.Update(r.Context(), updateRequest)

	webResponse := &dto.WebResponse{
		Message: "Success Update User",
		Data:    updateResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (c *UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c.UserService.Logout(r.Context())

	webResponse := &dto.WebResponse{
		Message: "Success Logout User",
		Data:    true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(webResponse)
	if err != nil {
		panic(err)
	}
}
