// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: table_campaign_acount.sql

package da_generated

import (
	"context"
)

const getCampaignAccountByCampaignId = `-- name: GetCampaignAccountByCampaignId :one
select id, account_id, campaign_id, created_at, updated_at
from campaign_account where campaign_id = ?
`

func (q *Queries) GetCampaignAccountByCampaignId(ctx context.Context, campaignID int64) (*CampaignAccount, error) {
	row := q.queryRow(ctx, q.getCampaignAccountByCampaignIdStmt, getCampaignAccountByCampaignId, campaignID)
	var i CampaignAccount
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.CampaignID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}