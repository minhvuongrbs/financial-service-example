// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: table_voucher_user.sql

package da_generated

import (
	"context"
)

const createVoucherUser = `-- name: CreateVoucherUser :execrows
INSERT INTO voucher_user(id, voucher_id, user_id, is_active)
VALUES (?, ?, ?, ?)
`

type CreateVoucherUserParams struct {
	ID        int64 `json:"id"`
	VoucherID int64 `json:"voucher_id"`
	UserID    int64 `json:"user_id"`
	IsActive  bool  `json:"is_active"`
}

func (q *Queries) CreateVoucherUser(ctx context.Context, arg *CreateVoucherUserParams) (int64, error) {
	result, err := q.exec(ctx, q.createVoucherUserStmt, createVoucherUser,
		arg.ID,
		arg.VoucherID,
		arg.UserID,
		arg.IsActive,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
