CREATE TABLE IF NOT EXISTS todos
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP    NOT NULL
);