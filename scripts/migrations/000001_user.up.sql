CREATE TABLE IF NOT EXISTS tb_user (
    id                  INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    , created_at        DATETIME NOT NULL DEFAULT NOW()
    , updated_at        DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()
    , deleted_at        DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00'
    , version           INT UNSIGNED NOT NULL DEFAULT 0
    , uuid              CHAR(36) NOT NULL
    , name              VARCHAR(50) NOT NULL
    , email             VARCHAR(50) NOT NULL
    , pass_encrypted    CHAR(60) NOT NULL

    , UNIQUE INDEX `uuid_uidx` (`uuid` ASC)
    , UNIQUE INDEX `email_uidx` (`email` ASC)
);