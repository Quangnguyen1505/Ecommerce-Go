-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_v1_goose (
  id   BIGSERIAL PRIMARY KEY,
  name  TEXT     NOT NULL,
  age   INT     NOT NULL,
  sex   INT     NOT NULL,
  created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_v1_goose;
-- +goose StatementEnd
