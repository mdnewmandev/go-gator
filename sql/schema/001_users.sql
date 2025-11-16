-- +goose Up
CREATE TABLE users(
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;