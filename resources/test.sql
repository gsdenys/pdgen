CREATE TABLE IF NOT EXISTS test (
    id serial PRIMARY KEY,
    name text NOT NULL
);

COMMENT ON TABLE test IS 'table for test propose';

COMMENT ON COLUMN test.id IS 'sequencial unique identifier';

COMMENT ON COLUMN test.name IS 'name of test';