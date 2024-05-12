package app

import (
	"context"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/voucher"
)

type DefineVoucherHandler struct {
	voucherRepo voucherRepo
}

type voucherRepo interface {
	UpdateVoucher(ctx context.Context, v voucher.Voucher) error
}

func NewDefineVoucherHandler(voucherRepo voucherRepo) DefineVoucherHandler {
	return DefineVoucherHandler{
		voucherRepo: voucherRepo,
	}
}

func (h DefineVoucherHandler) Handle(ctx context.Context) {
	// validate input
	// store db
}
