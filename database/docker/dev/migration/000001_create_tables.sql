DROP TABLE IF EXISTS `collections`;
CREATE TABLE `collections`(
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(20) NOT NULL UNIQUE,
  `order` INT,
  active BOOLEAN NOT NULL DEFAULT TRUE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `items`;
CREATE TABLE `items` (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(80) NOT NULL UNIQUE,
  search_keys TEXT,
  description TEXT,
  fk_label_origin INT,
  fk_collection INT,
  fk_label_color INT,
  fk_image_avatar INT,
  `order_in_collection` INT,
  `order` INT,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  FULLTEXT (`name`, `search_keys`, `description`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `images`;
CREATE TABLE `images`(
  id INT AUTO_INCREMENT PRIMARY KEY,
  link VARCHAR(1024) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `item_images`;
CREATE TABLE `item_images`(
  fk_item INT NOT NULL,
  fk_image INT NOT NULL UNIQUE,
  `order` INT,
  is_avatar BOOLEAN NOT NULL DEFAULT FALSE,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  PRIMARY KEY (`fk_item`, `fk_image`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `labels`;
CREATE TABLE `labels` (
  id INT AUTO_INCREMENT PRIMARY KEY,
  code VARCHAR(40) NULL,
  value VARCHAR(255),
  fk_label INT,
  active BOOLEAN NOT NULL DEFAULT TRUE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `item_attributes`;
CREATE TABLE `item_attributes`(
  fk_label INT NOT NULL,
  fk_item INT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  PRIMARY KEY (`fk_item`, `fk_label`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`(
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(40) NOT NULL,
  hashed_password VARCHAR(255) NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`(
  id INT AUTO_INCREMENT PRIMARY KEY,
  label VARCHAR(255) NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles`(
  fk_user INT,
  fk_role INT,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  PRIMARY KEY (`fk_user`, `fk_role`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

--
-- DROP TABLE IF EXISTS `permissions`;
-- CREATE TABLE `permissions`(
--   fk_role INT,
--   fk_label_resource VARCHAR(40),
--   permit enum('read', 'write', 'read-all', 'edit')
-- )

