package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaignuser"
)

type JoinCampaignHandler struct {
	campaignUserRepo campaignUserRepo
}

var ErrCampaignUserNotFound = errors.New("campaign user not found")

type campaignUserRepo interface {
	GetCampaignUser(ctx context.Context, req JoinCampaignRequest) (*campaignuser.CampaignUser, error)
	InsertCampaignUser(ctx context.Context, req JoinCampaignRequest) error
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
		return fmt.Errorf("repository get campaign user faield: %w", err)
	}

	err = h.campaignUserRepo.InsertCampaignUser(ctx, req)
	if err != nil {
		return fmt.Errorf("repo insert campaign user failed: %w", err)
	}
	return nil
}
