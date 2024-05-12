package service

import (
	"context"
	"errors"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaign"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaignuser"
)

var ErrCampaignUserNotFound = errors.New("campaign user not found")

type campaignUserRepo interface {
	GetCampaignUser(ctx context.Context, req JoinCampaignRequest) (*campaignuser.CampaignUser, error)
	InsertCampaignUser(ctx context.Context, req JoinCampaignRequest) error
}

type campaignRepo interface {
	GetCampaignById(ctx context.Context, campaignId int64) (campaign.Campaign, error)
}
