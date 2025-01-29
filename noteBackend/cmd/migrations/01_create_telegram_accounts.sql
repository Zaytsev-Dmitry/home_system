-- +goose Up
CREATE TABLE telegram_accounts (
                      id SERIAL PRIMARY KEY,
                      account_id integer,
                      name varchar(100),
                      link varchar(100),
                      description varchar(100)
);