-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS urls (
                                  id INTEGER PRIMARY KEY autoincrement,
                                  alias TEXT NOT NULL UNIQUE,
                                  url TEXT NOT NULL,
                                  created_at DATETIME,
                                  updated_at DATETIME
);
CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls
-- +goose StatementEnd
