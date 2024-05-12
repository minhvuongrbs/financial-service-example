-- store user join campaign
CREATE TABLE campaign_user
(
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id     BIGINT NOT NULL,
    campaign_id BIGINT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk__campaign_user__campaign_id__user_id (campaign_id, user_id)
);

-- name: GetCampaignUser :one
select *
from campaign_user where campaign_id = ? and user_id = ?;

-- name: CreateCampaignUser :execrows
INSERT INTO campaign_user(campaign_id, user_id)
VALUES (?, ?)