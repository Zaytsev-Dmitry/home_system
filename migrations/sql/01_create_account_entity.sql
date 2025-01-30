CREATE TABLE accounts
(
    id         SERIAL PRIMARY KEY,
    first_name varchar(100),
    last_name  varchar(100),
    login      varchar(100),
    email      varchar(100),
    type       varchar(10),
    telegram_id int
);
