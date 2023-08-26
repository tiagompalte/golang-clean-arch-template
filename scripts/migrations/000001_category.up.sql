CREATE TABLE IF NOT EXISTS tb_category (
    id              INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
    , created_at    DATETIME NOT NULL DEFAULT NOW()
    , updated_at    DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW()
    , deleted_at    DATETIME NOT NULL DEFAULT '0000-00-00'
    , slug          VARCHAR(50) NOT NULL
    , name          VARCHAR(50) NOT NULL

    , UNIQUE INDEX `slug_uidx` ( 
        `slug`          ASC
        , `deleted_at`  ASC
    )
);
