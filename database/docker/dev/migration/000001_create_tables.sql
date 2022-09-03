DROP TABLE IF EXISTS `collections`;
CREATE TABLE `collections`(
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(20) NOT NULL UNIQUE,
  `order` INT NOT NULL UNIQUE,
  active BOOLEAN NOT NULL DEFAULT TRUE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `items`;
CREATE TABLE `items` (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(20) NOT NULL UNIQUE,
  search_keys TEXT,
  description TEXT,
  `order` INT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  FULLTEXT (`name`, `search_keys`, `description`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `item_collection`;
CREATE TABLE `item_collection`(
  fk_item INT NOT NULL,
  fk_collection INT NOT NULL,
  `order` INT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  PRIMARY KEY (`fk_item`, `fk_collection`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `images`;
CREATE TABLE `images`(
  id INT AUTO_INCREMENT PRIMARY KEY,
  link VARCHAR(1024) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `item_image`;
CREATE TABLE `item_image`(
  fk_item INT NOT NULL,
  fk_image INT NOT NULL,
  `order` INT NOT NULL,
  is_avatar BOOLEAN NOT NULL DEFAULT FALSE,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  PRIMARY KEY (`fk_item`, `fk_image`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

-- DELIMITER $$
--
-- CREATE TRIGGER before_insert_allow_only_one_avatar_per_item
--   BEFORE INSERT ON stack_over_t FOR EACH ROW
-- BEGIN
--   IF (SELECT COUNT(*) FROM `item_image`
--       WHERE fk_item=NEW.`fk_item` AND item_image.is_avatar = TRUE) > 0 AND NEW.`is_avatar` = TRUE
--           THEN
--                SIGNAL SQLSTATE '45000'
--       SET MESSAGE_TEXT = 'Cannot add or update row: only one avatar allowed per item';
-- END IF;
-- END;
-- $$
--
-- CREATE TRIGGER before_update_allow_only_one_avatar_per_item
--   BEFORE UPDATE ON item_image FOR EACH ROW
-- BEGIN
--   IF (SELECT COUNT(*) FROM `item_image`
--       WHERE fk_item=NEW.`fk_item` AND item_image.is_avatar = TRUE) > 0 AND NEW.`is_avatar` = TRUE
--           THEN
--                SIGNAL SQLSTATE '45000'
--       SET MESSAGE_TEXT = 'Cannot add or update row: only one avatar allowed per item';
-- END IF;
-- END;
-- $$

DROP TABLE IF EXISTS `labels`;
CREATE TABLE `labels` (
  id INT AUTO_INCREMENT PRIMARY KEY,
  code VARCHAR(40) NULL UNIQUE,
  value VARCHAR(255),
  fk_label INT,
  active BOOLEAN NOT NULL DEFAULT TRUE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

DROP TABLE IF EXISTS `item_attribute`;
CREATE TABLE `item_attribute`(
  fk_label INT NOT NULL,
  fk_item INT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  PRIMARY KEY (`fk_item`, `fk_label`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

