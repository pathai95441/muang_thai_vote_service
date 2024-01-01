-- migrate:up
INSERT INTO `role` (`id`, `role_name`, `description`, `created_by`)
VALUES (1, 'Admin', 'Admin Role CRUD candidate & user', 'migration'),
       (2, 'Guest', 'Guest Role can only vote candidate', 'migration');

-- migrate:down
DELETE FROM `role` WHERE id = 1;
DELETE FROM `role` WHERE id = 2;
