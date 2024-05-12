package admin

import (
	"context"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaign"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/voucher"
)

type campaignRepo interface {
	DefineCampaign(ctx context.Context, c *campaign.Campaign) error
}

type voucherRepo interface {
	UpdateVoucher(ctx context.Context, v *voucher.Voucher) error
}
