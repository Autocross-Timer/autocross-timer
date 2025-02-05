CREATE TABLE IF NOT EXISTS `api_tokens`
(
 `api_token`    varchar(45) NOT NULL ,
 `user`         varchar(45) NOT NULL ,

PRIMARY KEY (`api_token`)
);