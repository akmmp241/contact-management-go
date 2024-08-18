package contracts

import (
	"contact-management-restful/models/dto"
	"context"
)

type UserService interface {
	Register(ctx context.Context, req dto.RegisterRequest) *dto.RegisterResponse
	Login(ctx context.Context, req dto.LoginRequest) *dto.LoginResponse
	Update(ctx context.Context, req dto.UpdateUserRequest) *dto.UpdateUserResponse
	Logout(ctx context.Context)
}

type ContactService interface {
	Create(ctx context.Context, req dto.CreateContactRequest) *dto.CreateContactResponse
	Search(ctx context.Context) *[]dto.SearchContactResponse
	Update(ctx context.Context, req dto.UpdateContactRequest, id int) *dto.UpdateContactResponse
	Get(ctx context.Context, id int) *dto.GetContactsResponse
	Delete(ctx context.Context, id int)
}

type AddressService interface {
	Create(ctx context.Context, req dto.AddressDTO, contactId int) *dto.AddressDTO
	Update(ctx context.Context, req dto.AddressDTO, contactId int, id int) *dto.AddressDTO
	Get(ctx context.Context, contactId int, id int) *dto.AddressDTO
	List(ctx context.Context, contactId int) *[]dto.AddressDTO
	Delete(ctx context.Context, contactId int, id int)
}
