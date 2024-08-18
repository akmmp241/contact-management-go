package helper

import (
	"contact-management-restful/models/domains"
	"contact-management-restful/models/dto"
)

func RegisterResponse(user *domains.User) *dto.RegisterResponse {
	return &dto.RegisterResponse{
		Name:     user.Name,
		Username: user.Username,
		Token:    user.Token,
	}
}
