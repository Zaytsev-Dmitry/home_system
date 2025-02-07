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
