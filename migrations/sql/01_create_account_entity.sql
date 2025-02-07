CREATE TABLE accounts
(
    id          SERIAL PRIMARY KEY,
    first_name  varchar(100),
    last_name   varchar(100),
    username    varchar(100),
    email       varchar(100),
    telegram_id int,
    keycloak_id text,
    is_active   boolean
);
grant all privileges on table accounts to auth_user;
grant all privileges on sequence accounts_id_seq to auth_user;
