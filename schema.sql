CREATE TABLE `urls` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `original` varchar(255) NOT NULL DEFAULT '',
  `tiny` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_tiny` (`tiny`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;