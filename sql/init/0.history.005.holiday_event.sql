CREATE TABLE `history`.`holiday_event` (
`id` INT NOT NULL AUTO_INCREMENT,
`date_time` DATETIME NULL,
`fact` VARCHAR(1000) NULL,
PRIMARY KEY (`id`),
INDEX `index2` (`date_time` ASC));
