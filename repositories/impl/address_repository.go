package impl

import (
	"contact-management-restful/models/domains"
	"context"
	"database/sql"
	"errors"
)

type AddressRepositoryImpl struct {
}

func NewAddressRepositoryImpl() *AddressRepositoryImpl {
	return &AddressRepositoryImpl{}
}

func (a AddressRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, address domains.Address) domains.Address {
	SQL := `INSERT INTO addresses (id, contact_id, street, city, province, country, postal_code) VALUES (NULL, ?, ?, ?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, SQL, address.ContactId, address.Street, address.City, address.Province, address.Country, address.Postcode)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	address.Id = int(lastInsertId)

	return address
}

func (a AddressRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, address domains.Address) domains.Address {
	SQL := `UPDATE addresses SET street = ?, city = ?, province = ?, country = ?, postal_code = ? WHERE id = ? AND contact_id = ?`
	_, err := tx.ExecContext(ctx, SQL, address.Street, address.City, address.Province, address.Country, address.Postcode, address.Id, address.ContactId)
	if err != nil {
		panic(err)
	}

	return address
}

func (a AddressRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, contactId int, id int) (domains.Address, error) {
	SQL := `SELECT id, contact_id, street, city, province, country, postal_code FROM addresses WHERE id = ? AND contact_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, id, contactId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return domains.Address{}, errors.New("address not found")
	}

	address := domains.Address{}
	err = rows.Scan(&address.Id, &address.ContactId, &address.Street, &address.City, &address.Province, &address.Country, &address.Postcode)
	if err != nil {
		panic(err)
	}

	return address, nil
}

func (a AddressRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, contactId int) []domains.Address {
	SQL := `SELECT id, contact_id, street, city, province, country, postal_code FROM addresses WHERE contact_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, contactId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var addresses []domains.Address
	for rows.Next() {
		address := domains.Address{}
		err := rows.Scan(&address.Id, &address.ContactId, &address.Street, &address.City, &address.Province, &address.Country, &address.Postcode)
		if err != nil {
			panic(err)
		}
		addresses = append(addresses, address)
	}

	return addresses
}

func (a AddressRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, id int) {
	SQL := `DELETE FROM addresses WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
}
