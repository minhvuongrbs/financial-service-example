CREATE TABLE `account`
(
    id            BIGINT                              NOT NULL AUTO_INCREMENT,
    username      VARCHAR(255) NULL UNIQUE,
    phone_number  VARCHAR(255) NULL UNIQUE COMMENT 'format 84914123423',
    email         VARCHAR(255) NULL UNIQUE,
    personal_info JSON                                NOT NULL DEFAULT JSON_OBJECT() COMMENT 'birthday,...',
    full_name     VARCHAR(1000)                       NOT NULL,
    password      VARCHAR(255)                        NOT NULL COMMENT 'hashed password',
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY pk__account (id)
);

-- name: CreateAccount :execresult
insert into account(username, phone_number, email, personal_info, full_name, password)
values (?, ?, ?, ?, ?, ?);

-- name: GetAccountByUsername :one
select *
from account
where username = ?;

-- name: GetAccountByPhoneNumber :one
select *
from account
where phone_number = ?;

-- name: GetAccountByEmail :one
select *
from account
where email = ?;
