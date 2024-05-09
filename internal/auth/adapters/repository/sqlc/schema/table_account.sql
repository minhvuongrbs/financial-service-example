CREATE TABLE `account`
(
    id            BIGINT                              NOT NULL AUTO_INCREMENT,
    username      VARCHAR(255) NULL UNIQUE,
    phone_number  VARCHAR(255) NULL UNIQUE,
    email         VARCHAR(255) NULL UNIQUE,
    personal_info JSON                                NOT NULL DEFAULT JSON_OBJECT() COMMENT 'birthday,...',
    full_name     VARCHAR(1000)                       NOT NULL,
    password      VARCHAR(255)                        NOT NULL COMMENT 'hashed password',
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY pk__account (id)
);