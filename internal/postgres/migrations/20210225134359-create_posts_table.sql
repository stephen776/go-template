
-- +migrate Up
CREATE TABLE IF NOT EXISTS posts (
  id       SERIAL        NOT NULL,
  title    VARCHAR(256)  NOT NULL,
  body     VARCHAR(5000) NOT NULL,
  user_id  INT           NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_post_user
    FOREIGN KEY (user_id)
      REFERENCES users(id)
);


-- +migrate Down
DROP TABLE IF EXISTS posts;