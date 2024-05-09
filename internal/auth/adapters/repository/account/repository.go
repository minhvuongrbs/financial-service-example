package account

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/minhvuongrbs/financial-service-example/internal/auth/adapters/repository/sqlc/da_generated"
	"github.com/minhvuongrbs/financial-service-example/internal/auth/entities/account"
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

func (r Repository) CreateAccount(ctx context.Context, a *account.Account) (*account.Account, error) {
	q := da_generated.New(r.db)
	marshalledPersonalInfo, err := json.Marshal(a.PersonalInfo)
	if err != nil {
		return nil, fmt.Errorf("marshal personal info: %w", err)
	}

	var marshaledUserName, marshaledPhone, marshaledEmail sql.NullString
	if a.Username != "" {
		marshaledUserName = sql.NullString{String: a.Username, Valid: true}
	}
	if a.PhoneNumber != 0 {
		marshaledPhone = sql.NullString{Valid: true, String: strconv.FormatInt(a.PhoneNumber, 10)}
	}
	if a.Email != "" {
		marshaledEmail = sql.NullString{Valid: true, String: a.Email}
	}

	result, err := q.CreateAccount(ctx, &da_generated.CreateAccountParams{
		Username:     marshaledUserName,
		PhoneNumber:  marshaledPhone,
		Email:        marshaledEmail,
		PersonalInfo: marshalledPersonalInfo,
		FullName:     a.FullName,
		Password:     a.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("da create account: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("rows affected: %w", err)
	}
	if rows != 1 {
		return nil, fmt.Errorf("invalid affected rows: %d", rows)
	}

	accountId, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("get last insert id failed: %w", err)
	}
	a.WithId(accountId)
	return a, nil
}

func (r Repository) GetAccountByIdentity(ctx context.Context, identity string,
	identityType account.IdentityType) (a *account.Account, err error) {
	q := da_generated.New(r.db)
	var daAccount *da_generated.Account

	switch identityType {
	case account.IdentityTypePhoneNumber:
		daAccount, err = q.GetAccountByPhoneNumber(ctx, sql.NullString{Valid: true, String: identity})
	case account.IdentityTypeEmail:
		daAccount, err = q.GetAccountByEmail(ctx, sql.NullString{Valid: true, String: identity})
	case account.IdentityTypeUsername:
		daAccount, err = q.GetAccountByUsername(ctx, sql.NullString{Valid: true, String: identity})
	default:
		return nil, fmt.Errorf("invalid identity type")
	}
	var unmarshalledPersonalInfo account.PersonalInfo
	err = json.Unmarshal(daAccount.PersonalInfo, &unmarshalledPersonalInfo)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	return &account.Account{
		Id:           daAccount.ID,
		Username:     daAccount.Username.String,
		PhoneNumber:  cast.ToInt64(daAccount.PhoneNumber),
		Email:        daAccount.Email.String,
		PersonalInfo: unmarshalledPersonalInfo,
		FullName:     daAccount.FullName,
		Password:     daAccount.Password,
	}, nil
}
