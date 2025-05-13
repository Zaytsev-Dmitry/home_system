CREATE TABLE category
(
    id        SERIAL PRIMARY KEY,
    name      varchar(1000) unique not null,
    is_active boolean default true not null
);


CREATE TABLE currency
(
    id        SERIAL PRIMARY KEY,
    name      varchar(1000) unique not null,
    is_active boolean default true not null
);

CREATE TABLE participant
(
    id               SERIAL PRIMARY KEY,
    keycloak_user_id varchar(1000) unique      not null,
    telegram_id      BIGINT unique             not null,
    display_name     varchar(1000),
    username         varchar(1000),
    is_active        boolean     default true  not null,
    created_at_utc   TIMESTAMPTZ DEFAULT now() not null
);

CREATE TABLE board
(
    id        SERIAL PRIMARY KEY,
    owner     bigint       NOT NULL references participant (id),
    name      varchar(100) NOT NULL,
    currency  bigint       NOT NULL references currency (id),
    is_active boolean      NOT NULL default true
);

CREATE TABLE board_participant
(
    board_id       INT                       NOT NULL REFERENCES board (id) ON DELETE CASCADE,
    participant_id INT                       NOT NULL REFERENCES participant (id) ON DELETE CASCADE,
    joined_at_utc  TIMESTAMPTZ DEFAULT now() NOT NULL,
    PRIMARY KEY (board_id, participant_id)
);

CREATE TABLE expense
(
    id             SERIAL PRIMARY KEY,
    title          varchar(1000)                      not null,
    participant_id bigint                             NOT NULL references participant (id),
    board_id       bigint                             NOT NULL references board (id),
    amount         decimal(10, 2) CHECK (amount >= 0) not null,
    created_at_utc TIMESTAMPTZ DEFAULT now()          not null,
    category_id    bigint                             NOT NULL references category (id)
);

CREATE TABLE expense_share
(
    expense_id     INT                                NOT NULL REFERENCES expense (id) ON DELETE CASCADE,
    participant_id INT                                NOT NULL REFERENCES participant (id) ON DELETE CASCADE,
    share_amount   DECIMAL(10, 2) CHECK (share_amount >= 0) NOT NULL,
    PRIMARY KEY (expense_id, participant_id)
);

CREATE INDEX idx_expense_participant_id ON expense (participant_id);
CREATE INDEX idx_expense_board_id ON expense (board_id);
CREATE INDEX idx_expense_category_id ON expense (category_id);
CREATE INDEX idx_expense_share_participant_id ON expense_share (participant_id);
CREATE INDEX idx_board_participant_participant_id ON board_participant (participant_id);