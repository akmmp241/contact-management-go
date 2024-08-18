package impl

import (
	"contact-management-restful/exception"
	"contact-management-restful/helper"
	"contact-management-restful/models/domains"
	"contact-management-restful/models/dto"
	"contact-management-restful/repositories/contracts"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type ContactServiceImpl struct {
	ContactRepository contracts.ContactRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewContactServiceImpl(contactRepository contracts.ContactRepository, DB *sql.DB, validate *validator.Validate) *ContactServiceImpl {
	return &ContactServiceImpl{ContactRepository: contactRepository, DB: DB, Validate: validate}
}

func (c ContactServiceImpl) Create(ctx context.Context, req dto.CreateContactRequest) *dto.CreateContactResponse {
	user := ctx.Value("user").(domains.User)
	err := c.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	tx, err := c.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contact := domains.Contact{
		UserId:    user.Id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}

	contact = c.ContactRepository.Save(ctx, tx, contact)

	contactResponse := dto.CreateContactResponse{
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
	}

	return &contactResponse
}

func (c ContactServiceImpl) Search(ctx context.Context) *[]dto.SearchContactResponse {
	tx, err := c.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contacts := c.ContactRepository.FindAll(ctx, tx)
	if len(contacts) < 1 {
		panic(exception.NewNotFound("contact not found"))
	}

	var searchContactResponse []dto.SearchContactResponse

	for _, contact := range contacts {
		searchContactResponse = append(searchContactResponse, dto.SearchContactResponse{
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Email:     contact.Email,
			Phone:     contact.Phone,
		})
	}

	return &searchContactResponse
}

func (c ContactServiceImpl) Update(ctx context.Context, req dto.UpdateContactRequest, id int) *dto.UpdateContactResponse {
	err := c.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	tx, err := c.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contact, err := c.ContactRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFound("contact not found"))
	}

	contact.FirstName = req.FirstName
	contact.LastName = req.LastName
	contact.Email = req.Email
	contact.Phone = req.Phone

	contact = c.ContactRepository.Update(ctx, tx, contact)

	contactResponse := dto.UpdateContactResponse{
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
	}

	return &contactResponse
}

func (c ContactServiceImpl) Get(ctx context.Context, id int) *dto.GetContactsResponse {
	tx, err := c.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contact, err := c.ContactRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFound("contact not found"))
	}

	contactResponse := dto.GetContactsResponse{
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
	}

	return &contactResponse
}

func (c ContactServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := c.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contact, err := c.ContactRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFound("contact not found"))
	}

	c.ContactRepository.DeleteById(ctx, tx, contact.Id)
}
