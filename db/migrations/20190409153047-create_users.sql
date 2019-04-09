
-- +migrate Up
create table users (
  id bigint unsigned auto_increment not null primary key,
  name varchar(255) not null,
  email varchar(255) not null,
  token varchar(255) not null unique,
  google_id varchar(255) not null unique,
  twitter_id varchar(255) not null unique,
  facebook_id varchar(255) not null unique,
  created_at datetime not null
);


-- +migrate Down
drop table if exists users;
