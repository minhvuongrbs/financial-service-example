CREATE TABLE account_voucher (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    account_id BIGINT NOT NULL,
    voucher_id BIGINT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- name: CreateAccountVoucher :execrows
INSERT INTO account_voucher(id, account_id, voucher_id, is_active)
VALUES (?, ?, ?, ?)
