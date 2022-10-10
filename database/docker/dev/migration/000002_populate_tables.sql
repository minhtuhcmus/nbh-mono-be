INSERT INTO `collections`(`name`, `order`)
VALUES
  ('rose', 1),
  ('daisy', 2),
  ('tulip', 3),
  ('dry', 4),
  ('extra', 5),
  ('hydrangea', 6),
  ('peony', 7),
  ('carnation', 8),
  ('gypsophila', 9),
  ('leaf', 10),
  ('pine', 11);

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES ('ITEM_ORIGIN', 'item-origin', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES
  ('item-origin', 'vietnam', @id),
  ('item-origin', 'ecuador', @id),
  ('item-origin', 'japan', @id),
  ('item-origin', 'south-africa', @id),
  ('item-origin', 'malaysia', @id),
  ('item-origin', 'israel', @id),
  ('item-origin', 'new-zealand', @id),
  ('item-origin', 'italy', @id),
  ('item-origin', 'united-states', @id),
  ('item-origin', 'australia', @id),
  ('item-origin', 'denmark', @id),
  ('item-origin', 'netherlands', @id);

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES ('ITEM_COLOR', 'item-color', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES
  ('item-color', 'red', @id),
  ('item-color', 'white', @id),
  ('item-color', 'brown' ,@id),
  ('item-color', 'dark-pink', @id),
  ('item-color', 'light-pink', @id),
  ('item-color', 'cream', @id),
  ('item-color', 'yellow', @id),
  ('item-color', 'orange', @id),
  ('item-color', 'green', @id),
  ('item-color', 'blue', @id),
  ('item-color', 'purple', @id);

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES ('ITEM_SIZE', 'item-size', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES
  ('item-size', '25cm', @id),
  ('item-size', '30cm', @id),
  ('item-size', '40cm' ,@id),
  ('item-size', '50cm', @id),
  ('item-size', '60cm', @id),
  ('item-size', '70cm', @id),
  ('item-size', '80cm', @id),
  ('item-size', '90cm', @id),
  ('item-size', '100cm', @id),
  ('item-size', '110cm', @id),
  ('item-size', '120cm', @id),
  ('item-size', '130cm', @id),
  ('item-size', '150cm', @id),
  ('item-size', '180cm', @id),
  ('item-size', '200cm', @id),
  ('item-size', '250cm', @id),
  ('item-size', '300cm', @id),
  ('item-size', '400cm', @id),
  ('item-size', '500cm', @id),
  ('item-size', '600cm', @id);

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES ('ITEM_AVAILABILITY', 'item-availability', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `value`, `fk_label`)
VALUES
  ('item-availability', 'january', @id),
  ('item-availability', 'february', @id),
  ('item-availability', 'march' ,@id),
  ('item-availability', 'april', @id),
  ('item-availability', 'may', @id),
  ('item-availability', 'june', @id),
  ('item-availability', 'july', @id),
  ('item-availability', 'august', @id),
  ('item-availability', 'september', @id),
  ('item-availability', 'october', @id),
  ('item-availability', 'november', @id),
  ('item-availability', 'december', @id);

INSERT INTO `users`(`username`, `hashed_password`)
VALUES ('admin', '$2a$14$yCNUd90EI3.fB3kOgvyS8Oht5097zjcfkmKtyvLNuaaAflqWDqpP6');

INSERT INTO `roles`(`label`)
VALUES ('admin'), ('customer');

INSERT INTO `user_role`(`fk_user`, `fk_role`)
VALUES (1, 1);
--
-- INSERT INTO `labels`(`code`, `value`, `fk_label`)
-- VALUES ('USER_PERMISSION', 'user-permission', NULL);
--
-- SET @id = LAST_INSERT_ID();
--
-- INSERT INTO `labels`(`code`, `value`, `fk_label`)
-- VALUES
--   ('user-permission', 'read', @id),
--   ('user-permission', 'write', @id),
--   ('user-permission', 'update' ,@id);
--
-- INSERT INTO `labels`(`code`, `value`, `fk_label`)
-- VALUES ('SYSTEM_RESOURCES', 'system-resources', NULL);
--
-- SET @id = LAST_INSERT_ID();
--
-- INSERT INTO `labels`(`code`, `value`, `fk_label`)
-- VALUES
--   ('system-resources', 'items', @id),
--   ('system-resources', 'roles', @id),
--   ('system-resources', 'users' ,@id),
--   ('system-resources', 'images' ,@id);

-- INSERT INTO `items`(`name`, `order`)
-- VALUES
--   ('celeste', 1),
--   ('sky-blue', 2),
--   ('blue', 3),
--   ('fancycake', 4),
--   ('rainbow', 5),
--   ('shyla', 6),
--   ('aquamarine', 7),
--   ('blue-&-white', 8),
--   ('orange-&-white', 9),
--   ('pink-&-white', 10),
--   ('purple-&-white', 11),
--   ('black-&-red', 12),
--   ('mondial', 13),
--   ('playa-blanca', 14),
--   ('quicksand', 15),
--   ('toffee', 16),
--   ('moody-blues', 17),
--   ('freedom', 18),
--   ('explorer', 19),
--   ('heart', 20),
--   ('kahala', 21),
--   ('freespirit', 22),
--   ('coral-reef', 23),
--   ('shimmer', 24),
--   ('country-sun', 25),
--   ('stardust', 26),
--   ('havana', 27),
--   ('pink-floyd', 28),
--   ('christa', 29),
--   ('full-monty', 30),
--   ('saga', 31),
--   ('moab', 32),
--   ('lemonade', 33),
--   ('elba', 34),
--   ('blue-moon', 35),
--   ('hot-majolica', 36),
--   ('brilliant-star-leo', 37),
--   ('girlie-folies', 38),
--   ('scarlet-mini', 39),
--   ('sweet-dream', 40),
--   ('rosa-cheer-girl-apricot', 41),
--   ('rosa-cheer-girl', 42),
--   ('rosa-sea-anemone', 43),
--   ('rosa-garnet-gem', 44),
--   ('rosa-mystic-sarah', 45);