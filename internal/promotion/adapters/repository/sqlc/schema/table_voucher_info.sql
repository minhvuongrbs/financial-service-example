CREATE TABLE voucher_info
(
    id              BIGINT AUTO_INCREMENT PRIMARY KEY,
    is_active       BOOL           NOT NULL DEFAULT FALSE,
    name           VARCHAR(1000)  NOT NULL,
    description     VARCHAR(1000)  NOT NULL DEFAULT '',
    metadata   JSON                        NOT NULL DEFAULT (JSON_OBJECT()) COMMENT 'metadata of voucher',
    created_at      TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP               DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);