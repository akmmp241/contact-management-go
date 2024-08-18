package impl

import (
	"contact-management-restful/models/domains"
	"context"
	"database/sql"
	"errors"
)

type ContactRepositoryImpl struct {
}

func NewContactRepositoryImpl() *ContactRepositoryImpl {
	return &ContactRepositoryImpl{}
}

func (c ContactRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, contact domains.Contact) domains.Contact {
	SQL := `INSERT INTO contacts (user_id, first_name, last_name, email, phone) VALUES (?, ?, ?, ?, ?)`
	_, err := tx.ExecContext(ctx, SQL, contact.UserId, contact.FirstName, contact.LastName, contact.Email, contact.Phone)
	if err != nil {
		panic(err)
	}

	return contact
}

func (c ContactRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domains.Contact {
	userId := ctx.Value("user").(domains.User).Id
	SQL := `SELECT id, user_id, first_name, last_name, email, phone FROM contacts WHERE user_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var contacts []domains.Contact
	for rows.Next() {
		var contact domains.Contact
		err = rows.Scan(&contact.Id, &contact.UserId, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
		if err != nil {
			panic(err)
		}
		contacts = append(contacts, contact)
	}

	return contacts
}

func (c ContactRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, contact domains.Contact) domains.Contact {
	SQL := `UPDATE contacts SET first_name = ?, last_name = ?, email = ?, phone = ? WHERE user_id = ? AND id = ?`
	_, err := tx.ExecContext(ctx, SQL, contact.FirstName, contact.LastName, contact.Email, contact.Phone, contact.UserId, contact.Id)
	if err != nil {
		panic(err)
	}

	return contact
}

func (c ContactRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domains.Contact, error) {
	userId := ctx.Value("user").(domains.User).Id
	SQL := `SELECT id, user_id, first_name, last_name, email, phone FROM contacts WHERE id = ? AND user_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, id, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var contact domains.Contact
	if !rows.Next() {
		return contact, errors.New("contact not found")
	}

	err = rows.Scan(&contact.Id, &contact.UserId, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
	if err != nil {
		panic(err)
	}

	return contact, nil
}

func (c ContactRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, id int) {
	userId := ctx.Value("user").(domains.User).Id
	SQL := `DELETE FROM contacts WHERE id = ? AND user_id = ?`
	_, err := tx.ExecContext(ctx, SQL, id, userId)
	if err != nil {
		panic(err)
	}
}
