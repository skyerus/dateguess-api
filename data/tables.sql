CREATE TABLE `history`.`session` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ip` VARCHAR(45) NULL,
  `date_time` DATETIME NULL,
  PRIMARY KEY (`id`),
  INDEX `index2` (`ip` ASC),
  INDEX `index3` (`date_time` ASC));
