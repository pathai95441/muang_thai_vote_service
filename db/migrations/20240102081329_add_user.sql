-- migrate:up
INSERT INTO `users` (`id`, `user_name`, `password`, `email`, `role_id`, `created_by`)
VALUES ('10fc104d-0571-47d6-9757-0defceba5613', 'admin', '$2a$10$W1B1i6KdPFFuqM2hiR06w.aEzMFtvzJybD5UV72.T3m5RDgRVbMKW', 'admin@test.com',1,'migration'),
       ('d2bcb9e8-0358-49b3-8b7c-dc4908d3065c', 'guest', '$2a$10$W1B1i6KdPFFuqM2hiR06w.aEzMFtvzJybD5UV72.T3m5RDgRVbMKW', 'guest@test.com',2,'migration');

-- migrate:down
DELETE FROM `users` WHERE id = '10fc104d-0571-47d6-9757-0defceba5613';
DELETE FROM `users` WHERE id = 'd2bcb9e8-0358-49b3-8b7c-dc4908d3065c';
