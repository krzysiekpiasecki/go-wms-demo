-- +goose Up
ALTER TABLE orders
ADD COLUMN comment TEXT NULL;


-- +goose Down
ALTER TABLE orders
DROP COLUMN comment;

