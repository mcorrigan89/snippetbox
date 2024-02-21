create table snippets
(
    id      serial
        primary key,
    title   varchar(100) not null,
    content text         not null,
    created timestamp    not null,
    expires timestamp    not null
);

alter table snippets
    owner to admin;

create index idx_snippets_created
    on snippets (created);

create table sessions
(
    token  char(43)     not null
        primary key,
    data   bytea        not null,
    expiry timestamp(6) not null
);

alter table sessions
    owner to admin;

create table users
(
    id              serial
        primary key,
    name            varchar(255) not null,
    email           varchar(255) not null
        constraint users_uc_email
            unique,
    hashed_password char(60)     not null,
    created         timestamp    not null
);

alter table users
    owner to admin;


INSERT INTO users (name, email, hashed_password, created) VALUES (
    'Alice Jones',
    'alice@example.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    '2022-01-01 09:18:24'
);