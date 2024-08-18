package middleware

import (
	"contact-management-restful/exception"
	"contact-management-restful/helper"
	"contact-management-restful/repositories/contracts"
	"context"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthMiddleware struct {
	UserRepository contracts.UserRepository
	DB             *sql.DB
}

func NewAuthMiddleware(userRepository contracts.UserRepository, DB *sql.DB) *AuthMiddleware {
	return &AuthMiddleware{UserRepository: userRepository, DB: DB}
}

func (h AuthMiddleware) ApiAuthMiddleware(next func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token := r.Header.Get("Authorization")
		authenticate := true

		if token == "" {
			authenticate = false
		}

		tx, err := h.DB.Begin()
		if err != nil {
			panic(err)
		}
		defer helper.CommitOrRollback(tx)

		user, err := h.UserRepository.FindByToken(r.Context(), tx, token)
		if err != nil {
			authenticate = false
		}
		ctx := context.WithValue(r.Context(), "user", user)

		if !authenticate {
			panic(exception.NewUnauthorized("Invalid token"))
		}

		next(w, r.WithContext(ctx), p)
	}
}
