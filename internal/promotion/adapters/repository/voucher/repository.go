package voucher

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/internal/promotion/adapters/repository/sqlc/da_generated"
	"github.com/minhvuongrbs/financial-service-example/internal/promotion/entities/voucher"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) UpdateVoucher(ctx context.Context, v *voucher.Voucher) error {
	q := da_generated.New(r.db)

	metadata, err := json.Marshal(v.Metadata)
	if err != nil {
		return fmt.Errorf("json marshal metadata failed: %w", err)
	}

	rows, err := q.UpdateVoucher(ctx, &da_generated.UpdateVoucherParams{
		ID:          v.Id,
		IsActive:    v.IsActive,
		Name:        v.Name,
		Description: v.Description,
		Metadata:    metadata,
	})
	if err != nil {
		return fmt.Errorf("da update voucher failed: %w", err)
	}
	if rows != 1 {
		return fmt.Errorf("invalid rows affected: %d", rows)
	}
	return nil
}
