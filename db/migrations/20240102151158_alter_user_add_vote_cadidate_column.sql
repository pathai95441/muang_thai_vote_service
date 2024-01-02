-- migrate:up
ALTER TABLE `users`
ADD COLUMN `vote_candidate_id` VARCHAR(40);

ALTER TABLE `users`
ADD CONSTRAINT `fk_user_vote_candidate`
FOREIGN KEY (`vote_candidate_id`)
REFERENCES `candidate`(`id`);

-- migrate:down
ALTER TABLE `users`
DROP COLUMN `vote_candidate_id`;
