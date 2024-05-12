-- store all campaigns
CREATE TABLE campaign
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    status     ENUM ('active', 'inactive') NOT NULL DEFAULT 'active',
    name       VARCHAR(1000)               NOT NULL,
    metadata   JSON                        NOT NULL DEFAULT (JSON_OBJECT()) COMMENT 'metadata of campaign',
    created_at TIMESTAMP                            DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- name: UpdateCampaign :execrows
insert into campaign (id, status, name, metadata)
VALUES (?, sqlc.arg(status), sqlc.arg(name), sqlc.arg(metadata))
ON DUPLICATE KEY UPDATE
    status = sqlc.arg(status),
    name = sqlc.arg(name),
    metadata = sqlc.arg(metadata);

-- name: GetCampaignById :one
select *
from campaign where id = ?;