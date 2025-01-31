--auth DB
CREATE TABLE profile
(
    id                SERIAL PRIMARY KEY,
    account_id        integer REFERENCES accounts (id),
    role              varchar(100),
    telegram_username varchar(100)
);

alter table accounts add column is_active boolean default true;
alter table accounts drop column login;
