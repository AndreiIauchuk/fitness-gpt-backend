-- +goose Up
CREATE TABLE test (
    id int primary key
);

-- +goose Down
DROP TABLE test;