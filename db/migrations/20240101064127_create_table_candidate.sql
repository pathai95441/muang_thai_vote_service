-- migrate:up
CREATE TABLE `candidate` (
  `id` VARCHAR(40) NOT NULL,
  `candidate_name` VARCHAR(255) NOT NULL,
  `candidate_description` VARCHAR(255) NOT NULL,
  `vote_score` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` VARCHAR(36) NOT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` VARCHAR(36) NULL,
  `deleted_at` TIMESTAMP NULL,
  `deleted_by` VARCHAR(36) NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin;

-- migrate:down
DROP TABLE `candidate`;


