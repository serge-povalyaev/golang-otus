-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS events
(
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title         TEXT NOT NULL,
    date_start     TIMESTAMP NOT NULL,
    date_finish   TIMESTAMP NOT NULL,
    description   TEXT,
    user_id       NUMERIC NOT NULL,
    notify_before NUMERIC
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events
-- +goose StatementEnd
