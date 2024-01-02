-- migrate:up
CREATE TABLE `users` (
  `id` VARCHAR(40) NOT NULL,
  `user_name` VARCHAR(255) UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `role_id` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` VARCHAR(36) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` VARCHAR(36) NULL,
  `deleted_at` TIMESTAMP NULL,
  `deleted_by` VARCHAR(36) NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY `fk___user_role` (`role_id`) REFERENCES `role`(`id`),
  INDEX `idx__user__user_name` (`user_name`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin;

-- migrate:down
DROP TABLE `users`;


