-- store account join campaign
CREATE TABLE campaign_account
(
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    account_id  BIGINT NOT NULL,
    campaign_id BIGINT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk__campaign_account__campaign_id__account_id(campaign_id, account_id)
);

-- name: GetCampaignAccountByCampaignId :one
select *
from campaign_account where campaign_id = ?;