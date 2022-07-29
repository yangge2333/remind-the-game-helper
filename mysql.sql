-- test.history definition

CREATE TABLE `history`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `user_id`   int(11) DEFAULT NULL,
    `tender_id` int(11) DEFAULT NULL,
    `data`      text     DEFAULT NULL,
    `update_at` datetime DEFAULT NULL,
    `create_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- test.tender definition

CREATE TABLE `tender`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `name`      varchar(100) DEFAULT NULL,
    `update_at` datetime     DEFAULT NULL,
    `create_at` datetime     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- test.`user` definition

CREATE TABLE `user`
(
    `id`        int(11) NOT NULL AUTO_INCREMENT,
    `username`  varchar(100) DEFAULT NULL,
    `password`  varchar(100) DEFAULT NULL,
    `update_at` datetime     DEFAULT NULL,
    `create_at` datetime     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;