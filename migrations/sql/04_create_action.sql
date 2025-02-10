CREATE TABLE user_action
(
    id                   SERIAL PRIMARY KEY,
    telegram_user_id     integer,
    need_user_action     boolean,
    last_sent_message_id integer,
    command_name         varchar(100),
    command_state        varchar(100)
);
grant all privileges on table user_action to telegram_bot_user;
grant all privileges on sequence user_action_id_seq to telegram_bot_user;