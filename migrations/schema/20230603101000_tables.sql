-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY unique,
    balance INTEGER CHECK (balance >= 0),
    created_dt TIMESTAMP DEFAULT current_timestamp
);
-- +goose Down
DROP TABLE users;