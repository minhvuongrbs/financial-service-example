CREATE TABLE `user`
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
    PRIMARY KEY pk__user (id)
);

-- name: CreateUser :execresult
insert into user(username, phone_number, email, personal_info, full_name, password)
values (?, ?, ?, ?, ?, ?);

-- name: GetUserByUsernameUsername :one
select *
from user
where username = ?;

-- name: GetUserByUsernamePhoneNumber :one
select *
from user
where phone_number = ?;

-- name: GetUserByUsernameEmail :one
select *
from user
where email = ?;
