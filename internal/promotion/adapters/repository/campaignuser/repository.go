package campaignuser

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/adapters/repository/sqlc/da_generated"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/app/service"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaignuser"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) InsertCampaignUser(ctx context.Context, req service.JoinCampaignRequest) error {
	q := da_generated.New(r.db)

	rows, err := q.CreateCampaignUser(ctx, &da_generated.CreateCampaignUserParams{
		CampaignID: req.CampaignId,
		UserID:     req.UserId,
	})
	if err != nil {
		return fmt.Errorf("create campaign user failed: %w", err)
	}
	if rows != 1 {
		return fmt.Errorf("invalid rows updated: %d", rows)
	}

	return nil
}

func (r Repository) GetCampaignUser(ctx context.Context,
	req service.JoinCampaignRequest) (*campaignuser.CampaignUser, error) {
	q := da_generated.New(r.db)

	resp, err := q.GetCampaignUser(ctx, &da_generated.GetCampaignUserParams{
		CampaignID: req.CampaignId,
		UserID:     req.UserId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, service.ErrCampaignUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get campaign user failed: %w", err)
	}
	return campaignuser.NewCampaignUser(resp.CampaignID, resp.UserID), nil
}
