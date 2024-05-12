package user_token

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/minhvuongrbs/financial-service-example/internal/auth/entities/user"
	"github.com/minhvuongrbs/financial-service-example/pkg/jwt"
)

type JWTToken struct {
	generator     *jwt.JwtRSAGenerator
	tokenDuration time.Duration
}

type Config struct {
	PrivateKeyPath string        `json:"private_key_path" mapstructure:"private_key_path"`
	TokenDuration  time.Duration `json:"token_duration" mapstructure:"token_duration"`
}

func NewJWTToken(cfg Config) (JWTToken, error) {
	generator, err := jwt.NewJwtRSAGeneratorFromFile(cfg.PrivateKeyPath)
	if err != nil {
		return JWTToken{}, fmt.Errorf("new jwt rsa generator from file: %w", err)
	}

	return JWTToken{
		generator:     generator,
		tokenDuration: cfg.TokenDuration,
	}, nil
}

func (j JWTToken) GenerateToken(m user.TokenData) (string, error) {
	m.GeneratedAt = time.Now().UnixMilli()
	m.ExpiredAt = time.Now().Add(j.tokenDuration).UnixMilli()

	bs, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("marshal user authenticate data: %w", err)
	}
	tokenData := map[string]interface{}{}
	if err = json.Unmarshal(bs, &tokenData); err != nil {
		return "", fmt.Errorf("unmarshal user authenticate data: %w", err)
	}
	token, err := j.generator.GenerateToken(tokenData)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}
	return token, nil
}
