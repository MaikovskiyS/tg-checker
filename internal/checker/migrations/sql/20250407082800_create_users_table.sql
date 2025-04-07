-- +goose Up
-- SQL in this section is executed when the migration is applied
CREATE TABLE IF NOT EXISTS users(
    id serial primary key,
    telegram_id bigint not null unique,
    channel_id bigint not null,
    created_at timestamp default now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back
DROP TABLE IF EXISTS users;
