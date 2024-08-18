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

type AddressServiceImpl struct {
	AddressRepository contracts.AddressRepository
	ContactRepository contracts.ContactRepository
	DB                *sql.DB
	validator         *validator.Validate
}

func NewAddressServiceImpl(addressRepository contracts.AddressRepository, contactRepository contracts.ContactRepository, DB *sql.DB, validator *validator.Validate) *AddressServiceImpl {
	return &AddressServiceImpl{AddressRepository: addressRepository, ContactRepository: contactRepository, DB: DB, validator: validator}
}

func (s AddressServiceImpl) Create(ctx context.Context, req dto.AddressDTO, contactId int) *dto.AddressDTO {
	err := s.validator.Struct(req)
	if err != nil {
		panic(err)
	}

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contact, err := s.ContactRepository.FindById(ctx, tx, contactId)
	if err != nil {
		panic(exception.NewNotFound("contact Not Found"))
	}

	address := domains.Address{
		ContactId: contact.Id,
		Street:    req.Street,
		City:      req.City,
		Province:  req.Province,
		Country:   req.Country,
		Postcode:  req.PostCode,
	}

	address = s.AddressRepository.Save(ctx, tx, address)
	req.Id = address.Id

	return &req
}

func (s AddressServiceImpl) Update(ctx context.Context, req dto.AddressDTO, contactId int, id int) *dto.AddressDTO {
	err := s.validator.Struct(req)
	if err != nil {
		panic(err)
	}

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contact, err := s.ContactRepository.FindById(ctx, tx, contactId)
	if err != nil {
		panic(exception.NewNotFound("contact Not Found"))
	}

	address, err := s.AddressRepository.FindById(ctx, tx, contact.Id, id)
	if err != nil {
		panic(exception.NewNotFound("address Not Found"))
	}

	address.Street = req.Street
	address.City = req.City
	address.Province = req.Province
	address.Country = req.Country
	address.Postcode = req.PostCode

	address = s.AddressRepository.Update(ctx, tx, address)
	req.Id = address.Id

	return &req
}

func (s AddressServiceImpl) Get(ctx context.Context, contactId int, id int) *dto.AddressDTO {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contact, err := s.ContactRepository.FindById(ctx, tx, contactId)
	if err != nil {
		panic(exception.NewNotFound("contact Not Found"))
	}

	address, err := s.AddressRepository.FindById(ctx, tx, contact.Id, id)
	if err != nil {
		panic(exception.NewNotFound("address Not Found"))
	}

	return &dto.AddressDTO{
		Id:       address.Id,
		Street:   address.Street,
		City:     address.City,
		Province: address.Province,
		Country:  address.Country,
		PostCode: address.Postcode,
	}
}

func (s AddressServiceImpl) List(ctx context.Context, contactId int) *[]dto.AddressDTO {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contact, err := s.ContactRepository.FindById(ctx, tx, contactId)
	if err != nil {
		panic(exception.NewNotFound("contact Not Found"))
	}

	addresses := s.AddressRepository.FindAll(ctx, tx, contact.Id)
	if len(addresses) < 1 {
		panic(exception.NewNotFound("address Not Found"))
	}

	var listAddresses []dto.AddressDTO
	for _, address := range addresses {
		listAddresses = append(listAddresses, dto.AddressDTO{
			Id:       address.Id,
			Street:   address.Street,
			City:     address.City,
			Province: address.Province,
			Country:  address.Country,
			PostCode: address.Postcode,
		})
	}

	return &listAddresses
}

func (s AddressServiceImpl) Delete(ctx context.Context, contactId int, id int) {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	contact, err := s.ContactRepository.FindById(ctx, tx, contactId)
	if err != nil {
		panic(exception.NewNotFound("contact Not Found"))
	}

	address, err := s.AddressRepository.FindById(ctx, tx, contact.Id, id)
	if err != nil {
		panic(exception.NewNotFound("address Not Found"))
	}

	s.AddressRepository.DeleteById(ctx, tx, address.Id)
}
