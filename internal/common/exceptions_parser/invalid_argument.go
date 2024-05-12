package exceptions_parser

import (
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/common/exceptions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func fromInvalidArgumentExceptions2GrpcStatus(ierr *exceptions.InvalidArgumentError) *status.Status {
	md := make(map[string]string)
	for k, v := range ierr.Metadata {
		md[k] = fmt.Sprintf("%v", v)
	}
	sts := status.New(codes.InvalidArgument, "Yêu cầu không hợp lệ.")
	return sts
}
