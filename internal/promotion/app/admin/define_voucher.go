package admin

import (
	"context"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/voucher"
)

type DefineVoucherHandler struct {
	voucherRepo voucherRepo
}

func NewDefineVoucherHandler(voucherRepo voucherRepo) DefineVoucherHandler {
	return DefineVoucherHandler{
		voucherRepo: voucherRepo,
	}
}

type DefineVoucherRequest struct {
	Name        string
	Description string
	IsActive    bool
	Metadata    voucher.Metadata
}

func (h DefineVoucherHandler) Handle(ctx context.Context, req DefineVoucherRequest) error {
	v := voucher.NewVoucher(0, req.Name, req.Description, req.IsActive, req.Metadata)
	err := h.voucherRepo.UpdateVoucher(ctx, v)
	if err != nil {
		return fmt.Errorf("repository update voucher failed: %w", err)
	}
	return nil
}
