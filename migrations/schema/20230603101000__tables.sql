-- +goose Up

CREATE TYPE operation_type AS ENUM ('accrual', 'redeem', 'send', 'receive');
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    balance DECIMAL DEFAULT 0.000 CHECK  (balance >= 0),
    created_dt TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE operations (
    id SERIAL PRIMARY KEY,
    operation_type operation_type,
    amount DECIMAL NOT NULL CHECK  (amount != 0),
    created_dt TIMESTAMP DEFAULT current_timestamp
);
CREATE TABLE user_operations (
    user_id INTEGER REFERENCES users (id),
    operation_id INTEGER REFERENCES operations(id),
    PRIMARY KEY (user_id, operation_id)
);
-- +goose Down
DROP TABLE user_operations;

DROP TABLE users;

DROP TABLE operations;

DROP TYPE operation_type;