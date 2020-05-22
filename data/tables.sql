CREATE TABLE `history`.`session` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ip` VARCHAR(45) NULL,
  `date_time` DATETIME NULL,
  PRIMARY KEY (`id`),
  INDEX `index2` (`ip` ASC),
  INDEX `index3` (`date_time` ASC));

  CREATE TABLE `history`.`historical_event` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `date_time` DATETIME NULL,
  `fact` VARCHAR(1000) NULL,
  PRIMARY KEY (`id`),
  INDEX `index2` (`date_time` ASC));

  CREATE TABLE `history`.`birth_event` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `date_time` DATETIME NULL,
  `fact` VARCHAR(1000) NULL,
  PRIMARY KEY (`id`),
  INDEX `index2` (`date_time` ASC));

CREATE TABLE `history`.`death_event` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `date_time` DATETIME NULL,
  `fact` VARCHAR(1000) NULL,
  PRIMARY KEY (`id`),
  INDEX `index2` (`date_time` ASC));

CREATE TABLE `history`.`holiday_event` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `date_time` DATETIME NULL,
  `fact` VARCHAR(1000) NULL,
  PRIMARY KEY (`id`),
  INDEX `index2` (`date_time` ASC));


