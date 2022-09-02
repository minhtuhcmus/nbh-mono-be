INSERT INTO `collections`(`name`, `order`)
VALUES
  ('ROSE', 1),
  ('DAISY', 2),
  ('TULIP', 3),
  ('DRY', 4),
  ('EXTRA', 5),
  ('HYDRANGEA', 6),
  ('PEONY', 7),
  ('CARNATION', 8),
  ('GYPSOPHILA', 9),
  ('LEAF', 10);

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES ('ITEM_ORIGIN', 'Item origin', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES
  ('VN', 'Vietnam', @id),
  ('ECU', 'Ecuador', @id),
  ('JP', 'Japan', @id),
  ('SA', 'South Africa', @id),
  ('ML', 'Malaysia', @id),
  ('IL', 'Israel', @id),
  ('NZ', 'New Zealand', @id),
  ('IT', 'Italy', @id),
  ('US', 'United States', @id),
  ('AU', 'Australia', @id),
  ('DK', 'Denmark', @id),
  ('NL', 'Netherlands', @id);

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES ('ITEM_COLOR', 'Item color', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES
  ('#FF0000', 'RED', @id),
  ('#FFFFFF', 'WHITE', @id),
  ('#964B00', 'BROWN' ,@id),
  ('#E06666', 'DARK_PINK', @id),
  ('#F4CCCC', 'LIGHT_PINK', @id),
  ('#FFFDD0', 'CREAM', @id),
  ('#FFFF00', 'YELLOW', @id),
  ('#FFA500', 'ORANGE', @id),
  ('#008000', 'GREEN', @id),
  ('#0000FF', 'BLUE', @id),
  ('#800080', 'PURPLE', @id);

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES ('ITEM_SIZE', 'Item size', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES
  ('25cm', '25cm', @id),
  ('30cm', '30cm', @id),
  ('40cm', '40cm' ,@id),
  ('50cm', '50cm', @id),
  ('60cm', '60cm', @id),
  ('70cm', '70cm', @id),
  ('80cm', '80cm', @id),
  ('90cm', '90cm', @id),
  ('100cm', '100cm', @id),
  ('110cm', '110cm', @id),
  ('120cm', '120cm', @id),
  ('130cm', '130cm', @id),
  ('150cm', '150cm', @id),
  ('180cm', '180cm', @id),
  ('200cm', '200cm', @id),
  ('250cm', '250cm', @id),
  ('300cm', '300cm', @id),
  ('400cm', '400cm', @id),
  ('500cm', '500cm', @id),
  ('600cm', '600cm', @id);

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES ('ITEM_AVAILABILITY', 'Item availability', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES
  ('JAN', 'January', @id),
  ('FEB', 'February', @id),
  ('MAR', 'March' ,@id),
  ('APR', 'April', @id),
  ('MAY', 'May', @id),
  ('JUN', 'June', @id),
  ('JUL', 'July', @id),
  ('AUG', 'August', @id),
  ('SEP', 'September', @id),
  ('OCT', 'October', @id),
  ('NOV', 'November', @id),
  ('DEC', 'December', @id);
