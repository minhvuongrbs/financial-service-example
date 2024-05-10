package app

import (
	"context"
	"fmt"
	"time"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaign"
)

type DefineCampaignHandler struct {
	campaignRepo campaignRepo
}

type DefineCampaignRequest struct {
	CampaignKey  string
	CampaignName string
}

func (h DefineCampaignHandler) Handle(ctx context.Context, req DefineCampaignRequest) error {
	c, err := campaign.NewCampaign(req.CampaignKey, req.CampaignName, campaign.StatusActive, time.Time{}, time.Time{})
	if err != nil {
		return fmt.Errorf("invalid campaign info: %w", err)
	}
	err = h.campaignRepo.CreateCampaign(ctx, c)
	if err != nil {
		return fmt.Errorf("repository create campaign: %w", err)
	}
	return nil
}
