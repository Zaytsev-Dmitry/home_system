CREATE TABLE accounts
(
    id                 SERIAL PRIMARY KEY,
    first_name         varchar(100),
    last_name          varchar(100),
    username           varchar(100),
    telegram_user_name varchar(200),
    email              varchar(100),
    telegram_id        int,
    keycloak_id        text,
    is_active          boolean
);

CREATE TABLE profile
(
    id         SERIAL PRIMARY KEY,
    account_id integer REFERENCES accounts (id),
    role       varchar(100),
    is_active  boolean,
    username   varchar(100)
);

grant all privileges on table profile to auth_user;
grant all privileges on sequence profile_id_seq to auth_user;
grant all privileges on table accounts to auth_user;
grant all privileges on sequence accounts_id_seq to auth_user;