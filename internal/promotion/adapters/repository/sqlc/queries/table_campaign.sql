-- store all campaigns
CREATE TABLE campaign
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    campaign_key VARCHAR(500)  NOT NULL UNIQUE KEY,
    status       ENUM('active', 'inactive') NOT NULL DEFAULT 'active',
    name         VARCHAR(1000) NOT NULL,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- name: GetCampaignByKey :one
select *
from campaign where campaign_key = ?;

-- name: InsertCampaign :execrows
insert into campaign (campaign_key, status, name, start_at, end_at)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateCampaignById :execrows
update campaign set status = ? and name = ? and start_at = ? and end_at = ?
where id = ?