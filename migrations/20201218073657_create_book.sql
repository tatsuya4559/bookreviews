-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE book (
  isbn CHAR(14) NOT NULL,
  title VARCHAR(255) NOT NULL,
  publisher VARCHAR(255) NOT NULL,
  PRIMARY KEY (isbn)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE book;
