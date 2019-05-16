
-- +migrate Up
alter table cards add index index_user_id(user_id);
alter table cards add index index_question_time(question_time);


-- +migrate Down
alter table cards drop index index_user_id;
alter table cards drop index index_question_time;
