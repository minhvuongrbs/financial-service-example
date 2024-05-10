// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package da_generated

import (
	"context"
)

type Querier interface {
	CreateAccountVoucher(ctx context.Context, arg *CreateAccountVoucherParams) (int64, error)
	GetCampaignAccountByCampaignId(ctx context.Context, campaignID int64) (*CampaignAccount, error)
	GetCampaignByKey(ctx context.Context, campaignKey string) (*Campaign, error)
	InsertCampaign(ctx context.Context, arg *InsertCampaignParams) (int64, error)
	InsertVoucher(ctx context.Context, arg *InsertVoucherParams) (int64, error)
	UpdateCampaignById(ctx context.Context, arg *UpdateCampaignByIdParams) (int64, error)
}

var _ Querier = (*Queries)(nil)