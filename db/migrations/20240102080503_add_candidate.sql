-- migrate:up
INSERT INTO `candidate` (`id`, `candidate_name`, `candidate_description`, `vote_score`, `created_by`)
VALUES (
    '1c4caf7b-b1da-4e72-8979-d79ac4ccd66c',
    'Pathai Laoatthphong',
    'Software Engineer at KKP DIME',
    0,
    'migration'
),
(
    '51d27f82-ca57-441c-9ba0-c5a6806ef71e',
    'Mark Zuckerberg',
    'Zuckerberg briefly attended Harvard University, where he launched Facebook in February 2004 with his roommates Eduardo Saverin',
    0,
    'migration'
),
(
    'dd8eefa4-71d1-40a6-9f26-92a415b1d687',
    'Bill Gate',
    'Later in his career and since leaving day-to-day operations at Microsoft in 2008, Gates has pursued other business and philanthropic endeavors.',
    0,
    'migration'
),
(
    '01fefd92-acc8-47f8-bc12-0173a76beb4a',
    'Elon Musk',
    'Musk has expressed views that have made him a polarizing figure.[7][8][9] He has been criticized for making unscientific and misleading statements',
    0,
    'migration'
),
(
    'd835c1d8-ab28-4922-8e96-5d1ac8b77019',
    'Steve Jobs',
    'In 1997, Jobs returned to Apple as CEO after the company''s acquisition of NeXT. He was largely responsible for reviving Apple,',
    0,
    'migration'
);


-- migrate:down
DELETE FROM `candidate` WHERE id = '1c4caf7b-b1da-4e72-8979-d79ac4ccd66c';
DELETE FROM `candidate` WHERE id = '51d27f82-ca57-441c-9ba0-c5a6806ef71e';
DELETE FROM `candidate` WHERE id = 'dd8eefa4-71d1-40a6-9f26-92a415b1d687';
DELETE FROM `candidate` WHERE id = '01fefd92-acc8-47f8-bc12-0173a76beb4a';
DELETE FROM `candidate` WHERE id = 'd835c1d8-ab28-4922-8e96-5d1ac8b77019';
