-- migrate:up
CREATE TABLE `vote_config` (
  `id` INT NOT NULL,
  `vote_close` TINYINT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` VARCHAR(36) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` VARCHAR(36) NULL,
  `deleted_at` TIMESTAMP NULL,
  `deleted_by` VARCHAR(36) NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin;

-- migrate:down
DROP TABLE `vote_config`;


