// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: table_account.sql

package da_generated

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createAccount = `-- name: CreateAccount :execresult
insert into account(username, phone_number, email, personal_info, full_name, password)
values (?, ?, ?, ?, ?, ?)
`

type CreateAccountParams struct {
	Username     sql.NullString  `json:"username"`
	PhoneNumber  sql.NullString  `json:"phone_number"`
	Email        sql.NullString  `json:"email"`
	PersonalInfo json.RawMessage `json:"personal_info"`
	FullName     string          `json:"full_name"`
	Password     string          `json:"password"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg *CreateAccountParams) (sql.Result, error) {
	return q.exec(ctx, q.createAccountStmt, createAccount,
		arg.Username,
		arg.PhoneNumber,
		arg.Email,
		arg.PersonalInfo,
		arg.FullName,
		arg.Password,
	)
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
select id, username, phone_number, email, personal_info, full_name, password, created_at, updated_at
from account
where email = ?
`

func (q *Queries) GetAccountByEmail(ctx context.Context, email sql.NullString) (*Account, error) {
	row := q.queryRow(ctx, q.getAccountByEmailStmt, getAccountByEmail, email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PhoneNumber,
		&i.Email,
		&i.PersonalInfo,
		&i.FullName,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getAccountByPhoneNumber = `-- name: GetAccountByPhoneNumber :one
select id, username, phone_number, email, personal_info, full_name, password, created_at, updated_at
from account
where phone_number = ?
`

func (q *Queries) GetAccountByPhoneNumber(ctx context.Context, phoneNumber sql.NullString) (*Account, error) {
	row := q.queryRow(ctx, q.getAccountByPhoneNumberStmt, getAccountByPhoneNumber, phoneNumber)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PhoneNumber,
		&i.Email,
		&i.PersonalInfo,
		&i.FullName,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getAccountByUsername = `-- name: GetAccountByUsername :one
select id, username, phone_number, email, personal_info, full_name, password, created_at, updated_at
from account
where username = ?
`

func (q *Queries) GetAccountByUsername(ctx context.Context, username sql.NullString) (*Account, error) {
	row := q.queryRow(ctx, q.getAccountByUsernameStmt, getAccountByUsername, username)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PhoneNumber,
		&i.Email,
		&i.PersonalInfo,
		&i.FullName,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}