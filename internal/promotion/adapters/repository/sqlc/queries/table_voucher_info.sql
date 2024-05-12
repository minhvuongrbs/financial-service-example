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

-- name: UpdateVoucher :execrows
insert into voucher_info (id, is_active, name, description, metadata)
VALUES (?, sqlc.arg(is_active), sqlc.arg(name), sqlc.arg(description), sqlc.arg(metadata))
ON DUPLICATE KEY UPDATE
     is_active = sqlc.arg(is_active),
     name = sqlc.arg(name),
     description = sqlc.arg(description),
     metadata = sqlc.arg(metadata);