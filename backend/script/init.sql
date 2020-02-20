CREATE SCHEMA `stock` DEFAULT CHARACTER SET utf8;

USE stock

CREATE TABLE `stock`.`user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `reward` int(11) DEFAULT 0,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

