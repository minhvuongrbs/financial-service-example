// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: table_user.sql

package da_generated

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createUser = `-- name: CreateUser :execresult
insert into user(username, phone_number, email, personal_info, full_name, password)
values (?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	Username     sql.NullString  `json:"username"`
	PhoneNumber  sql.NullString  `json:"phone_number"`
	Email        sql.NullString  `json:"email"`
	PersonalInfo json.RawMessage `json:"personal_info"`
	FullName     string          `json:"full_name"`
	Password     string          `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) (sql.Result, error) {
	return q.exec(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.PhoneNumber,
		arg.Email,
		arg.PersonalInfo,
		arg.FullName,
		arg.Password,
	)
}

const getUserByUsernameEmail = `-- name: GetUserByUsernameEmail :one
select id, username, phone_number, email, personal_info, full_name, password, created_at, updated_at
from user
where email = ?
`

func (q *Queries) GetUserByUsernameEmail(ctx context.Context, email sql.NullString) (*User, error) {
	row := q.queryRow(ctx, q.getUserByUsernameEmailStmt, getUserByUsernameEmail, email)
	var i User
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

const getUserByUsernamePhoneNumber = `-- name: GetUserByUsernamePhoneNumber :one
select id, username, phone_number, email, personal_info, full_name, password, created_at, updated_at
from user
where phone_number = ?
`

func (q *Queries) GetUserByUsernamePhoneNumber(ctx context.Context, phoneNumber sql.NullString) (*User, error) {
	row := q.queryRow(ctx, q.getUserByUsernamePhoneNumberStmt, getUserByUsernamePhoneNumber, phoneNumber)
	var i User
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

const getUserByUsernameUsername = `-- name: GetUserByUsernameUsername :one
select id, username, phone_number, email, personal_info, full_name, password, created_at, updated_at
from user
where username = ?
`

func (q *Queries) GetUserByUsernameUsername(ctx context.Context, username sql.NullString) (*User, error) {
	row := q.queryRow(ctx, q.getUserByUsernameUsernameStmt, getUserByUsernameUsername, username)
	var i User
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