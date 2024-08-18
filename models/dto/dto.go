package dto

type WebResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"error"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse RegisterResponse

type CurrentResponse RegisterResponse

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
}

type UpdateUserResponse RegisterResponse

type CreateContactRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,number"`
}

type CreateContactResponse CreateContactRequest

type SearchContactResponse CreateContactRequest

type UpdateContactRequest CreateContactRequest

type UpdateContactResponse CreateContactRequest

type GetContactsResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type AddressDTO struct {
	Id       int    `json:"id"`
	Street   string `json:"street" validate:"required"`
	City     string `json:"city" validate:"required"`
	Province string `json:"province" validate:"required"`
	Country  string `json:"country" validate:"required"`
	PostCode string `json:"postal_code" validate:"required,number"`
}
