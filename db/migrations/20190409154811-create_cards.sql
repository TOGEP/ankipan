
-- +migrate Up
create table cards (
  id bigint unsigned auto_increment not null primary key,
  user_id bigint unsigned not null,
  problem_statement text not null,
  answer_text text not null,
  memo text not null,
  question_time datetime not null default NOW()
);

-- +migrate Down
drop table if exists cards;
