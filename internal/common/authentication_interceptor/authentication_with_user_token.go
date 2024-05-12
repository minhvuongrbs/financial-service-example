package authentication_interceptor

import (
	"context"
	"fmt"
	"strings"

	"github.com/minhvuongrbs/financial-service-example/internal/common/exceptions_parser"
	"github.com/minhvuongrbs/financial-service-example/internal/common/validate_user_token"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/authenticate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthenticationWithUserToken struct {
	validateUserToken         validate_user_token.ValidateUserToken
	methodsByPassAuthenticate map[string]struct{}
}

func NewAuthenticationWithUserToken(v validate_user_token.ValidateUserToken, methodsByPassAuthenticate []string) AuthenticationWithUserToken {
	m := make(map[string]struct{})
	for _, method := range methodsByPassAuthenticate {
		m[method] = struct{}{}
	}
	return AuthenticationWithUserToken{
		validateUserToken:         v,
		methodsByPassAuthenticate: m,
	}
}

func (a *AuthenticationWithUserToken) ValidateToken(ctx context.Context, userToken string) (authenticate.TokenData, error) {
	data, err := a.validateUserToken.Handle(ctx, userToken)
	if err != nil {
		return authenticate.TokenData{}, fmt.Errorf("validate user token: %w", err)
	}
	return data, nil
}

const (
	authenticationHeader = "authorization"
)

type ApiRequestForUser interface {
	GetUserId() int64
}

func (a *AuthenticationWithUserToken) UserTokenAuthenticationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if _, ok := a.methodsByPassAuthenticate[info.FullMethod]; ok {
			return handler(ctx, req)
		}
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			sts := status.New(codes.PermissionDenied, "thông tin xác thực không hợp lệ.")
			return nil, sts.Err()
		}
		tokens := md.Get(authenticationHeader)
		if len(tokens) == 0 {
			sts := status.New(codes.PermissionDenied, "thông tin xác thực không hợp lệ.")
			return nil, sts.Err()
		}
		token := strings.Replace(tokens[0], "Bearer ", "", 1)
		data, err := a.ValidateToken(ctx, token)
		if err != nil {
			sts := exceptions_parser.Err2GrpcStatus(err)
			return nil, sts.Err()
		}
		if data.UserId == 0 {
			sts := status.New(codes.PermissionDenied, "thông tin xác thực không hợp lệ.")
			return nil, sts.Err()
		}
		if ureq, ok := req.(ApiRequestForUser); ok {
			if ureq.GetUserId() != data.UserId {
				sts := status.New(codes.PermissionDenied, "thông tin xác thực không hợp lệ.")
				return nil, sts.Err()
			}
		}

		ctx = authenticate.ContextWithAuthentication(ctx, data)
		return handler(ctx, req)
	}
}
