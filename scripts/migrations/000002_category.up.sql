CREATE TABLE IF NOT EXISTS tb_category (
    id              INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    , created_at    DATETIME NOT NULL DEFAULT NOW()
    , updated_at    DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()
    , deleted_at    DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00'
    , slug          VARCHAR(50) NOT NULL
    , name          VARCHAR(50) NOT NULL
    , user_id       INT UNSIGNED NOT NULL

    , UNIQUE INDEX `slug_uidx` ( 
        `slug`          ASC
        , `user_id`     ASC
        , `deleted_at`  ASC
    )
    , INDEX `user_uidx` (
        `user_id`       ASC
        , `deleted_at`  ASC
    )
    , CONSTRAINT fk_category_user_id
        FOREIGN KEY (user_id)
        REFERENCES tb_user (id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
);