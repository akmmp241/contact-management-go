package contracts

import (
	"contact-management-restful/models/domains"
	"context"
	"database/sql"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domains.User) domains.User
	Update(ctx context.Context, tx *sql.Tx, user domains.User) domains.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domains.User, error)
	FindByToken(ctx context.Context, tx *sql.Tx, token string) (domains.User, error)
}

type ContactRepository interface {
	Save(ctx context.Context, tx *sql.Tx, contact domains.Contact) domains.Contact
	Update(ctx context.Context, tx *sql.Tx, contact domains.Contact) domains.Contact
	FindById(ctx context.Context, tx *sql.Tx, id int) (domains.Contact, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domains.Contact
	DeleteById(ctx context.Context, tx *sql.Tx, id int)
}

type AddressRepository interface {
	Save(ctx context.Context, tx *sql.Tx, address domains.Address) domains.Address
	Update(ctx context.Context, tx *sql.Tx, address domains.Address) domains.Address
	FindById(ctx context.Context, tx *sql.Tx, contactId int, id int) (domains.Address, error)
	FindAll(ctx context.Context, tx *sql.Tx, contactId int) []domains.Address
	DeleteById(ctx context.Context, tx *sql.Tx, id int)
}
