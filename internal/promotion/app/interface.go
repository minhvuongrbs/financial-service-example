package app

import (
	"context"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaign"
)

type campaignRepo interface {
	DefineCampaign(ctx context.Context, c *campaign.Campaign) error
}
