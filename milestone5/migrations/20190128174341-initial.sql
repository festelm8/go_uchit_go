
-- +migrate Up
CREATE TABLE `users` (
	`id` INT(5) NOT NULL AUTO_INCREMENT,
	`phone` VARCHAR(25) DEFAULT NULL,
	`password` VARCHAR(255) NOT NULL,
	`name` VARCHAR(255) DEFAULT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB;

-- +migrate Down
DROP TABLE users;