CREATE TABLE voucher_user
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    voucher_id BIGINT  NOT NULL,
    user_id BIGINT  NOT NULL,
    is_active  BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- name: CreateVoucherUser :execrows
INSERT INTO voucher_user(id, voucher_id, user_id, is_active)
VALUES (?, ?, ?, ?)
