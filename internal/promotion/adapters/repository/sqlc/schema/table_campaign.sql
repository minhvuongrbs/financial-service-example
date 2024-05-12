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