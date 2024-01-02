-- migrate:up
INSERT INTO `vote_config` (`id`, `vote_close`, `created_by`)
VALUES (1,0,'migration');

-- migrate:down
DELETE FROM `vote_config` WHERE id = 1;
