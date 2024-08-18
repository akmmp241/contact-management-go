package impl

import (
	"contact-management-restful/models/domains"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domains.User) domains.User {
	query := `INSERT INTO users (ID, USERNAME, PASSWORD, NAME, TOKEN) VALUES (null, ?, ?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, user.Username, user.Password, user.Name, user.Token)
	if err != nil {
		panic(err)
	}

	return user
}

func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domains.User) domains.User {
	query := `UPDATE users SET username = ?, name = ?, password = ?, token = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, query, user.Username, user.Name, user.Password, user.Token, user.Id)
	if err != nil {
		panic(err)
	}

	return user
}

func (u *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domains.User, error) {
	query := `SELECT id, username, password, name, token FROM users WHERE username = ?`
	rows, err := tx.QueryContext(ctx, query, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var user domains.User
	if !rows.Next() {
		return user, errors.New("user not found")
	}

	err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.Token)
	if err != nil {
		panic(err)
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindByToken(ctx context.Context, tx *sql.Tx, token string) (domains.User, error) {
	query := `SELECT id, username, password, name, token FROM users WHERE token = ?`
	rows, err := tx.QueryContext(ctx, query, token)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var user domains.User
	if !rows.Next() {
		return user, errors.New("user not found")
	}

	err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.Token)
	if err != nil {
		panic(err)
	}

	return user, nil
}
