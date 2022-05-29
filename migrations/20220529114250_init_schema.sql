-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts
(
  id INTEGER NOT NULL,
  balance NUMERIC NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO accounts(id, balance)
VALUES 
(1, 100.0),
(2, 35.0),
(3, 90.0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
-- +goose StatementEnd
