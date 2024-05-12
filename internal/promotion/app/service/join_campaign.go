package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaign"
)

type JoinCampaignHandler struct {
	campaignUserRepo campaignUserRepo
	campaignRepo     campaignRepo
}

func NewJoinCampaignHandler(campaignUserRepo campaignUserRepo) JoinCampaignHandler {
	return JoinCampaignHandler{campaignUserRepo: campaignUserRepo}
}

type JoinCampaignRequest struct {
	CampaignId int64
	UserId     int64
}

func (h JoinCampaignHandler) Handle(ctx context.Context, req JoinCampaignRequest) error {
	_, err := h.campaignUserRepo.GetCampaignUser(ctx, req)
	if err == nil {
		// user joined campaign => return not error
		return nil
	}
	if err != nil && !errors.Is(err, ErrCampaignUserNotFound) {
		return fmt.Errorf("repository get campaign user failed: %w", err)
	}
	c, err := h.campaignRepo.GetCampaignById(ctx, req.CampaignId)
	if err != nil {
		return fmt.Errorf("repo get campaign by id failed: %w", err)
	}
	if c.Status != campaign.StatusActive {
		return fmt.Errorf("campaign is not active, status: %v", c.Status)
	}

	err = h.campaignUserRepo.InsertCampaignUser(ctx, req)
	if err != nil {
		return fmt.Errorf("repo insert campaign user failed: %w", err)
	}
	return nil
}
