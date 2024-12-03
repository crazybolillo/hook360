-- migrate:up
CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    payload JSONB NOT NULL
)

-- migrate:down
