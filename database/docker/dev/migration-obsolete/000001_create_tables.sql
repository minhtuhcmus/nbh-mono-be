-- DROP TABLE IF EXISTS `users`;
-- CREATE TABLE `users` (
--   id INT AUTO_INCREMENT PRIMARY KEY,
--   username VARCHAR(255) NOT NULL UNIQUE,
--   display_name VARCHAR(255) NULL,
--   password_hashed VARCHAR(64) NOT NULL,
--   active BOOLEAN NOT NULL DEFAULT TRUE,
--   phone_number VARCHAR(20) NOT NULL UNIQUE,
--   fk_label_country_code INT NOT NULL,
--   fk_role INT NOT NULL,
--   fk_label_status INT NOT NULL,
--   created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
--   updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
--   KEY `idx_display_name` (`display_name`),
--   KEY `idx_username` (`username`),
--   KEY `idx_phone_number` (`phone_number`)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

-- DROP TABLE IF EXISTS `customers`;
-- CREATE TABLE `customers` (
--   id INT AUTO_INCREMENT PRIMARY KEY,
--   fk_user INT NOT NULL UNIQUE,
--   email VARCHAR(255) NOT NULL UNIQUE,
--   active BOOLEAN NOT NULL DEFAULT TRUE,
--   shop_name VARCHAR(255) NULL,
--   address TEXT NOT NULL,
--   created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
--   updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

-- DROP TABLE IF EXISTS `roles`;
-- CREATE TABLE `roles` (
--   id INT AUTO_INCREMENT PRIMARY KEY,
--   name VARCHAR(20) NOT NULL UNIQUE,
--   description VARCHAR(255),
--   active BOOLEAN NOT NULL DEFAULT TRUE
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `labels`;
CREATE TABLE `labels` (
  id INT AUTO_INCREMENT PRIMARY KEY,
  code VARCHAR(40) NULL UNIQUE,
  description VARCHAR(255),
  fk_label INT,
  active BOOLEAN NOT NULL DEFAULT TRUE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `items`;
CREATE TABLE `items` (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(20) NOT NULL UNIQUE,
  search_keys VARCHAR(1024) NULL UNIQUE,
--   display_price INT NOT NULL,
--   retail_price INT NULL,
--   fk_label_item_currency INT NOT NULL,
--   fk_label_item_origin INT NOT NULL,
--   fk_label_item_type INT NOT NULL,
--   fk_packing INT NOT NULL,
--   fk_sell_block INT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  KEY `idx_name` (`name`),
  KEY `idx_display_name` (`display_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

-- DROP TABLE IF EXISTS `item_color`;
-- CREATE TABLE `item_color`(
--   fk_item INT NOT NULL,
--   fk_color INT NOT NULL,
--   active BOOLEAN NOT NULL DEFAULT TRUE,
--   PRIMARY KEY (`fk_item`, `fk_color`)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;
--
-- DROP TABLE IF EXISTS `packings`;
-- CREATE TABLE `packings`(
--   id INT AUTO_INCREMENT PRIMARY KEY,
--   number INT NOT NULL,
--   fk_label_item_unit INT NOT NULL,
--   fk_label_packing_style VARCHAR(40) NOT NULL,
--   active BOOLEAN NOT NULL DEFAULT TRUE
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;
--
-- DROP TABLE IF EXISTS `sell_blocks`;
-- CREATE TABLE `sell_blocks`(
--   id INT AUTO_INCREMENT PRIMARY KEY,
--   number INT NOT NULL,
--   fk_label_item_unit VARCHAR(40) NOT NULL,
--   active BOOLEAN NOT NULL DEFAULT TRUE
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;
--
-- DROP TABLE IF EXISTS `permissions`;
-- CREATE TABLE `permissions`(
--   id INT AUTO_INCREMENT PRIMARY KEY,
--   fk_user INT,
--   fk_role INT,
--   fk_label_resource INT NOT NULL,
--   active BOOLEAN NOT NULL DEFAULT TRUE
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `images`;
CREATE TABLE `images`(
  id INT AUTO_INCREMENT PRIMARY KEY,
  link VARCHAR(1024) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `item_images`;
CREATE TABLE `item_images`(
  fk_item INT NOT NULL,
  fk_image INT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  PRIMARY KEY (`fk_item`, `fk_image`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `item_attributes`;
CREATE TABLE `item_attributes`(
  fk_label INT NOT NULL,
  fk_item INT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  PRIMARY KEY (`fk_item`, `fk_label`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

