package campaign

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/adapters/repository/sqlc/da_generated"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/campaign"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) DefineCampaign(ctx context.Context, c *campaign.Campaign) error {
	q := da_generated.New(r.db)

	daStatus, err := toDACampaignStatus(c.Status)
	if err != nil {
		return fmt.Errorf("to DA campaign status: %w", err)
	}
	metadata, err := json.Marshal(c.Metadata)
	if err != nil {
		return fmt.Errorf("json marshal vourchers failed: %w", err)
	}

	rows, err := q.UpdateCampaign(ctx, &da_generated.UpdateCampaignParams{
		ID:       c.Id,
		Status:   daStatus,
		Name:     c.Name,
		Metadata: metadata,
	})
	if err != nil {
		return fmt.Errorf("da insert campaign failed: %w", err)
	}
	if rows != 1 {
		return fmt.Errorf("invalid rows affected: %d", rows)
	}
	return nil
}

func toDACampaignStatus(status campaign.Status) (da_generated.CampaignStatus, error) {
	switch status {
	case campaign.StatusActive:
		return da_generated.CampaignStatusActive, nil
	case campaign.StatusInactive:
		return da_generated.CampaignStatusInactive, nil
	default:
		return "", fmt.Errorf("invalid status: %v", status)
	}
}
