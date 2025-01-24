-- +goose Up
CREATE TABLE notes (
                      id SERIAL PRIMARY KEY,
                      account_id integer,
                      name varchar(100),
                      link varchar(100)
);