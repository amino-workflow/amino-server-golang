CREATE TABLE IF NOT EXISTS workflows (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    flow VARCHAR,
    version INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);