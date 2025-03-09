CREATE TABLE notes
(
    id          SERIAL PRIMARY KEY,
    account_id  integer,
    telegram_id integer,
    name        varchar(100),
    link        varchar(100),
    description varchar(100)
);
grant all privileges on table notes to notes_user;
grant all privileges on sequence notes_id_seq to notes_user;