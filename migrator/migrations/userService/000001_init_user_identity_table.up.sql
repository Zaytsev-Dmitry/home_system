CREATE TABLE user_identity_links
(
    id               SERIAL PRIMARY KEY,
    keycloak_user_id UUID         NOT NULL,
    telegram_user_id BIGINT       NOT NULL,
    email            varchar(500) NOT NULL,
    created_at       TIMESTAMP DEFAULT now(),
    UNIQUE (keycloak_user_id),
    UNIQUE (telegram_user_id),
    UNIQUE (email)
);