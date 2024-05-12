package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/minhvuongrbs/financial-service-example/internal/auth/adapters/repository/sqlc/da_generated"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/entities/user"
	"github.com/spf13/cast"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) CreateUser(ctx context.Context, u *user.User) (*user.User, error) {
	q := da_generated.New(r.db)
	marshalledPersonalInfo, err := json.Marshal(u.PersonalInfo)
	if err != nil {
		return nil, fmt.Errorf("marshal personal info: %w", err)
	}

	var marshaledUserName, marshaledPhone, marshaledEmail sql.NullString
	if u.Username != "" {
		marshaledUserName = sql.NullString{String: u.Username, Valid: true}
	}
	if u.PhoneNumber != 0 {
		marshaledPhone = sql.NullString{Valid: true, String: strconv.FormatInt(u.PhoneNumber, 10)}
	}
	if u.Email != "" {
		marshaledEmail = sql.NullString{Valid: true, String: u.Email}
	}

	result, err := q.CreateUser(ctx, &da_generated.CreateUserParams{
		Username:     marshaledUserName,
		PhoneNumber:  marshaledPhone,
		Email:        marshaledEmail,
		PersonalInfo: marshalledPersonalInfo,
		FullName:     u.FullName,
		Password:     u.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("da create user: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("rows affected: %w", err)
	}
	if rows != 1 {
		return nil, fmt.Errorf("invalid affected rows: %d", rows)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("get last insert id failed: %w", err)
	}
	u.WithId(userId)
	return u, nil
}

func (r Repository) GetUserByUsernameIdentity(ctx context.Context, identity string,
	identityType user.IdentityType) (a *user.User, err error) {
	q := da_generated.New(r.db)
	var daUser *da_generated.User

	switch identityType {
	case user.IdentityTypePhoneNumber:
		daUser, err = q.GetUserByUsernamePhoneNumber(ctx, sql.NullString{Valid: true, String: identity})
	case user.IdentityTypeEmail:
		daUser, err = q.GetUserByUsernameEmail(ctx, sql.NullString{Valid: true, String: identity})
	case user.IdentityTypeUsername:
		daUser, err = q.GetUserByUsernameUsername(ctx, sql.NullString{Valid: true, String: identity})
	default:
		return nil, fmt.Errorf("invalid identity type")
	}
	var unmarshalledPersonalInfo user.PersonalInfo
	err = json.Unmarshal(daUser.PersonalInfo, &unmarshalledPersonalInfo)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	return &user.User{
		Id:           daUser.ID,
		Username:     daUser.Username.String,
		PhoneNumber:  cast.ToInt64(daUser.PhoneNumber),
		Email:        daUser.Email.String,
		PersonalInfo: unmarshalledPersonalInfo,
		FullName:     daUser.FullName,
		Password:     daUser.Password,
	}, nil
}
