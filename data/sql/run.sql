CREATE TABLE IF NOT EXISTS `runs`
(
 `event_id`     varchar(45) NOT NULL ,
 `run_number`   int(255) NOT NULL ,
 `car_number`   varchar(45) NOT NULL ,
 `raw_time`     varchar(45) NOT NULL ,
 `pax_time`     varchar(45) NOT NULL ,
 `car_class`    varchar(45) NOT NULL ,
 `driver_name`  varchar(45) NOT NULL ,
 `cones`        int(255) NULL ,
 `is_dnf`       tinyint(1) NOT NULL DEFAULT 0,
 `gets_rerun`   tinyint(1) NOT NULL DEFAULT 0,
 `last_updated` int(255) NOT NULL ,
 `created`      int(255) NOT NULL,

PRIMARY KEY (`event_id`, `run_number`)
);