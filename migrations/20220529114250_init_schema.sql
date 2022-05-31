-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts
(
  id SERIAL NOT NULL,
  balance NUMERIC NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO accounts(balance)
VALUES 
(100.0),
(35.0),
(90.0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
-- +goose StatementEnd
