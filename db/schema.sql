CREATE SCHEMA IF NOT EXISTS `arch-lab3`;
USE `arch-lab3`;

DROP TABLE IF EXISTS `arch-lab3`.`balancers`;
CREATE TABLE IF NOT EXISTS `arch-lab3`.`balancers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `usedMachines` VARCHAR(255) NOT NULL,
  `totalMachinesCount` INT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

insert into `balancers` (usedMachines, totalMachinesCount) values ("[1,2,3]", 3);
insert into `balancers` (usedMachines, totalMachinesCount) values ("[4,5,6]", 3);

DROP TABLE IF EXISTS `arch-lab3`.`machines`;
CREATE TABLE IF NOT EXISTS `arch-lab3`.`machines` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `os` VARCHAR(45) NOT NULL,
  `processor` VARCHAR(45) NOT NULL,
  `workingStatus` VARCHAR(45) NOT NULL,
  `connectedTo` INT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

insert into `machines` (os, processor, workingStatus, connectedTo) values ("linux debian", "intel core i3", "yes", 1);
insert into `machines` (os, processor, workingStatus, connectedTo) values ("mac os monterey", "intel core i5", "yes", 1);
insert into `machines` (os, processor, workingStatus, connectedTo) values ("linux mint", "intel core i7", "yes", 1);
insert into `machines` (os, processor, workingStatus, connectedTo) values ("linux ubuntu", "intel core i5", "yes", 2);
insert into `machines` (os, processor, workingStatus, connectedTo) values ("linux debian", "amd ryzen 3", "yes", 2);
insert into `machines` (os, processor, workingStatus, connectedTo) values ("windows 10", "amd ryzen 5", "yes", 2);
