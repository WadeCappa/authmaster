
create table users (
    username varchar(256),
    hash varchar(256),
    user_id bigint,
    
    primary key (user_id),
    unique (username)
);

create index user_lookup on users (username, hash);

create sequence user_ids start 101;

create table tokens (
    user_id bigint,
    token varchar(256),
    expire_time timestamptz,

    foreign key (user_id) references users(user_id),
    unique(token),
    primary key (user_id)
);

create index token_lookup on tokens (token);