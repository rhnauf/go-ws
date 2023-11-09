-- +goose Up
CREATE TABLE messages (
   id UUID PRIMARY KEY,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   username VARCHAR(100) NOT NULL,
   recipient VARCHAR(100) NOT NULL,
    message TEXT NOT NULL
);

-- +goose Down
DROP TABLE messages;
