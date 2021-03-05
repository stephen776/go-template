
-- +migrate Up
ALTER TABLE users
RENAME COLUMN lasst_name TO last_name;

-- +migrate Down
