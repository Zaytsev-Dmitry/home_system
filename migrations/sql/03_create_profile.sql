--auth DB
CREATE TABLE profile
(
    id         SERIAL PRIMARY KEY,
    account_id integer REFERENCES accounts (id),
    role       varchar(100),
    is_active  boolean,
    username   varchar(100)
);
