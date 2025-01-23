-- +goose Up
CREATE TABLE account (
                      id int NOT NULL PRIMARY KEY,
                      first_name varchar(100),
                      second_name varchar(100),
                      login varchar(100)
);