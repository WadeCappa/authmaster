create database authmaster_db;

\c authmaster_db

create table if not exists users (
    username varchar(256),
    hash varchar(256),
    user_id bigint,
    
    primary key (user_id),
    unique (username)
);

create sequence if not exists user_ids start 101;

create table if not exists tokens (
    user_id bigint,
    token varchar(256),
    expire_time timestamptz,

    foreign key (user_id) references users(user_id) on delete cascade,
    unique(token),
    primary key (user_id, token)
);

create index if not exists token_lookup on tokens (token);