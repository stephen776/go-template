
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id         SERIAL       NOT NULL,
  email      VARCHAR(256) NOT NULL,
  first_name VARCHAR(50)  NOT NULL,
  lasst_name VARCHAR(50)  NOT NULL,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS users;