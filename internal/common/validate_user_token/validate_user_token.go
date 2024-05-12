package validate_user_token

import (
	"context"
	"time"

	"github.com/minhvuongrbs/financial-service-example/internal/common/exceptions"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/authenticate"
)

type ValidateUserToken struct {
	validator tokenValidator
}

func NewValidateUserToken(v tokenValidator) (ValidateUserToken, error) {
	if v == nil {
		return ValidateUserToken{}, exceptions.NewInvalidArgumentError("validator", "validator must not nil", nil)
	}
	return ValidateUserToken{
		validator: v,
	}, nil
}

//go:generate mockgen --source=./validate_user_token.go --destination=./mocks.go --package=validate_user_token .

type tokenValidator interface {
	ValidateToken(token string) (authenticate.TokenData, error)
}

func (v ValidateUserToken) Handle(_ context.Context, token string) (authenticate.TokenData, error) {
	if token == "" {
		return authenticate.TokenData{}, exceptions.NewPreconditionError(exceptions.PreconditionReasonInvalidToken, exceptions.SubjectAuthentication, "token must not empty", nil)
	}
	u, err := v.validator.ValidateToken(token)
	if err != nil {
		return authenticate.TokenData{}, exceptions.NewPreconditionError(
			exceptions.PreconditionReasonInvalidToken,
			exceptions.SubjectAuthentication,
			"token invalid",
			map[string]interface{}{
				"validate": err,
			},
		)
	}

	if u.ExpiredAt < time.Now().UnixMilli() {
		return authenticate.TokenData{}, exceptions.NewPreconditionError(
			exceptions.PreconditionReasonTokenExpired,
			exceptions.SubjectAuthentication,
			"token expired",
			map[string]interface{}{
				"expired_at": u.ExpiredAt,
			},
		)
	}

	return u, nil
}
