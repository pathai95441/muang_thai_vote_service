-- migrate:up
CREATE TABLE `vote_history` (
  `id` VARCHAR(40) NOT NULL,
  `candidate_id` VARCHAR(40) NOT NULL,
  `user_id` VARCHAR(40) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` VARCHAR(36) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` VARCHAR(36) NULL,
  `deleted_at` TIMESTAMP NULL,
  `deleted_by` VARCHAR(36) NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY `fk___candidate_id` (`candidate_id`) REFERENCES `candidate`(`id`),
  FOREIGN KEY `fk___user_id` (`user_id`) REFERENCES `users`(`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin;

-- migrate:down
DROP TABLE `vote_history`;