INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES ('USER_STATUS', 'Status of a user', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES
       ('ACTIVE', 'active', @id);

SET @active_id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES
       ('SUSPEND', 'suspend', @id),
       ('QUEUE', 'queue', @id);

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES ('ITEM_TYPE', 'Item type', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES
  ('ROSE', 'Hoa hồng', @id),
  ('DAISY', 'Cúc mẫu đơn', @id),
  ('TULIP', 'Hoa Tulip', @id),
  ('DRY', 'Đồ khô', @id),
  ('EXTRA', 'Hoa phụ', @id),
  ('HYDRANGEA', 'Cẩm tú cầu', @id);

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES ('COUNTRY_CODE', 'Country code', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES
  ('VN', '+84', @id);

SET @vn_country_code = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES ('ITEM_CURRENCY', 'Item currency', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES
  ('VNĐ', 'đ', @id);

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES ('ITEM_ORIGIN', 'Item origin', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES
  ('VNM', 'Việt Nam', @id),
  ('ECU', 'Ecuador', @id),
  ('JP', 'Nhật Bản', @id),
  ('SA', 'Nam Phi', @id),
  ('ML', 'Malaysia', @id),
  ('NL', 'Hà Lan', @id);

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES ('PACKING_UNIT', 'Packing unit', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES
  ('STEM', 'cành', @id),
  ('BUNDLE', 'bó', @id),
  ('KG', 'kg', @id),
  ('BOX', 'hộp', @id),
  ('PACKAGE', 'thùng', @id);

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES ('RESOURCE', 'Resource', NULL);

SET @id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES ('ADMIN_RESOURCE', 'admin_resource', @id);

SET @admin_resource_label_id = LAST_INSERT_ID();

INSERT INTO `labels`(`code`, `description`, `fk_label`)
VALUES
  ('MUTATION_SIGNUP', 'mutation_signup', @id),
  ('MUTATION_SIGNIN', 'mutation_signin', @id),
  ('MUTATION_CREATE_USER', 'mutation_creat_user', @id),
  ('MUTATION_UPDATE_USER', 'mutation_update_user', @id),
  ('MUTATION_UPDATE_ME', 'mutation_update_me', @id),
  ('MUTATION_CREATE_ITEM', 'mutation_creat_item', @id),
  ('MUTATION_UPDATE_ITEM', 'mutation_update_item', @id),
  ('MUTATION_CREATE_ROLE', 'mutation_creat_role', @id),
  ('MUTATION_UPDATE_ROLE', 'mutation_update_role', @id),
  ('MUTATION_CREATE_LABEL', 'mutation_creat_label', @id),
  ('MUTATION_UPDATE_LABEL', 'mutation_update_label', @id),
  ('MUTATION_CREATE_PERMISSION', 'mutation_creat_permission', @id),
  ('MUTATION_UPDATE_PERMISSION', 'mutation_update_permission', @id),
  ('QUERY_GET_ME', 'query_get_me', @id),
  ('QUERY_GET_USER', 'query_get_user', @id),
  ('QUERY_GET_ITEM', 'query_get_item', @id),
  ('QUERY_GET_LABEL', 'query_get_label', @id),
  ('QUERY_GET_ROLE', 'query_get_role', @id),
  ('QUERY_GET_PERMISSION', 'query_get_permission', @id),
  ('QUERY_LIST_USER', 'query_list_user', @id),
  ('QUERY_LIST_ITEM', 'query_list_item', @id),
  ('QUERY_LIST_LABEL', 'query_list_label', @id),
  ('QUERY_LIST_ROLE', 'query_list_role', @id),
  ('QUERY_LIST_PERMISSION', 'query_list_permission', @id);

INSERT INTO `roles`(`description`, `description`)
VALUES ('ADMIN', 'have access to every module');

SET @admin_role_id = LAST_INSERT_ID();

INSERT INTO `roles`(`description`, `description`)
VALUES
  ('ACCOUNTANT', 'accountant_role'),
  ('MANAGER', 'manager_role'),
  ('DELIVERY', 'delivery_role'),
  ('SALE', 'sale_role'),
  ('STAFF', 'staff_role');

INSERT INTO `users`(
  `username`,
  `password_hashed`,
  `display_name`,
  `phone_number`,
  `fk_label_country_code`,
  `fk_role`,
  `fk_label_status`
) VALUES (
  'admin',
  '$2a$14$AMqLzju6Tzj.N2E8UvONGe7HNcTPfPYuWlPHw7ADr7znO3AnZdBki',
	'ADMIN',
  'N/A',
  @vn_country_code,
  @admin_role_id,
  @active_id
);

INSERT INTO `permissions`(
  `fk_user`,
  `fk_role`,
  `fk_label_resource`
) VALUES (
  1,
  @admin_role_id,
  @admin_resource_label_id
)
       





