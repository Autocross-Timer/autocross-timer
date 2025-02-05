CREATE TABLE IF NOT EXISTS `events`
(
 `event_id`       varchar(45) NOT NULL ,
 `club_name`      varchar(45) NOT NULL ,
 `event_location` varchar(45) NOT NULL ,
 `event_date`     date NOT NULL ,
 `event_number`   int(255) NOT NULL ,

PRIMARY KEY (`event_id`)
);