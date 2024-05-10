CREATE TABLE voucher
(
    id              BIGINT AUTO_INCREMENT PRIMARY KEY,
    voucher_type    ENUM('fixed_amount', 'percentage_discount') NOT NULL,
    amount          DECIMAL(10, 2) NOT NULL,
    is_active       BOOL           NOT NULL DEFAULT FALSE,
    expiration_time DATETIME       NOT NULL,
    app_id          VARCHAR(255)   NOT NULL COMMENT '* will apply to all app_id',
    created_at      TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP               DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);