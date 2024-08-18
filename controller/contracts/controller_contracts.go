package contracts

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Current(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type ContactController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	List(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type AddressController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	List(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
