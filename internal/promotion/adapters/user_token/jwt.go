package user_token

import (
	"encoding/json"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/authenticate"
	"github.com/minhvuongrbs/financial-service-example/pkg/jwt"
)

type JWTToken struct {
	validator *jwt.JwtRSAValiator
}

type Config struct {
	PublicKeyPath string `json:"public_key_path" mapstructure:"public_key_path"`
}

func NewJWTToken(cfg Config) (JWTToken, error) {
	validator, err := jwt.NewJwtRSAValidatorFromFile(cfg.PublicKeyPath)
	if err != nil {
		return JWTToken{}, fmt.Errorf("new jwt rsa validator from file: %w", err)
	}

	return JWTToken{
		validator: validator,
	}, nil
}

func (j JWTToken) ValidateToken(token string) (authenticate.TokenData, error) {
	tokenData, err := j.validator.ValidateToken(token)
	if err != nil {
		return authenticate.TokenData{}, fmt.Errorf("validate token: %w", err)
	}
	bs, err := json.Marshal(tokenData)
	if err != nil {
		return authenticate.TokenData{}, fmt.Errorf("marshal token data: %w", err)
	}
	var u authenticate.TokenData
	if err := json.Unmarshal(bs, &u); err != nil {
		return authenticate.TokenData{}, fmt.Errorf("unmarshal token data: %w", err)
	}
	return u, nil
}
