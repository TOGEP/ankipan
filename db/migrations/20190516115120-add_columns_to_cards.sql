
-- +migrate Up
alter table cards add solved_count int unsigned not null;

-- +migrate Down
alter table cards drop column solved_count;
