-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE book (
  isbn VARCHAR(20) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  publisher VARCHAR(255) NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE book;
