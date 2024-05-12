package app

import (
	"context"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaign"
)

type DefineCampaignHandler struct {
	campaignRepo campaignRepo
}

func NewDefineCampaignHandler(campaignRepo campaignRepo) DefineCampaignHandler {
	return DefineCampaignHandler{
		campaignRepo: campaignRepo,
	}
}

type DefineCampaignRequest struct {
	Name     string
	status   campaign.Status
	Metadata campaign.Metadata
}

func (h DefineCampaignHandler) Handle(ctx context.Context, req DefineCampaignRequest) error {
	// todo: validate voucher
	c, err := campaign.NewCampaign(0, req.Name, campaign.StatusActive, req.Metadata)

	if err != nil {
		return fmt.Errorf("invalid campaign info: %w", err)
	}
	err = h.campaignRepo.DefineCampaign(ctx, c)
	if err != nil {
		return fmt.Errorf("repository create campaign: %w", err)
	}
	return nil
}
