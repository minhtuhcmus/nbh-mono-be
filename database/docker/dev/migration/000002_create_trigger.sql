CREATE TRIGGER `before_item_insert`
  BEFORE INSERT ON `items`
  FOR EACH ROW
  SET NEW.`order` = (SELECT COUNT(*)+1 FROM `items` WHERE `items`.`active` = true);

CREATE TRIGGER `before_collection_insert`
  BEFORE INSERT ON `collections`
  FOR EACH ROW
  SET NEW.`order` = (SELECT COUNT(*)+1 FROM `collections` WHERE `collections`.`active` = true);

CREATE TRIGGER `before_item_collection_insert`
  BEFORE INSERT ON `item_collections`
  FOR EACH ROW
  SET NEW.`order` = (SELECT COUNT(*)+1 FROM `item_collections` WHERE `item_collections`.`fk_collection` = NEW.`fk_collection` AND `item_collections`.`active` = true);

CREATE TRIGGER `before_item_image_insert`
  BEFORE INSERT ON `item_images`
  FOR EACH ROW
  SET NEW.`order` = (SELECT COUNT(*)+1 FROM `item_images` WHERE `item_images`.`fk_item` = NEW.`fk_item` AND `item_images`.`active` = true);
