CREATE TABLE IF NOT EXISTS tb_task (
    id              INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    , created_at    DATETIME NOT NULL DEFAULT NOW()
    , updated_at    DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()
    , deleted_at    DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00'
    , uuid          CHAR(36) NOT NULL
    , name          VARCHAR(50) NOT NULL
    , description   TEXT NOT NULL
    , done          BOOLEAN NOT NULL DEFAULT FALSE
    , category_id   INT UNSIGNED NOT NULL
    , user_id       INT UNSIGNED NOT NULL

    , UNIQUE INDEX `uuid_uidx` (`uuid` ASC)
    , INDEX `user_uidx` (
        `user_id`       ASC
        , `deleted_at`  ASC
    )
    , CONSTRAINT fk_task_category_id
        FOREIGN KEY (category_id)
        REFERENCES tb_category (id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
    , CONSTRAINT fk_task_user_id
        FOREIGN KEY (user_id)
        REFERENCES tb_user (id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
);