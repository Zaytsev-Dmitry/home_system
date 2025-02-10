CREATE TABLE user_action
(
    id                   SERIAL PRIMARY KEY,
    telegram_user_id     integer,
    last_action          varchar(50),
    last_requirement     varchar(50),
    last_sent_message_id integer
);
grant all privileges on table user_action to telegram_bot_user;
grant all privileges on sequence user_action_id_seq to telegram_bot_user;