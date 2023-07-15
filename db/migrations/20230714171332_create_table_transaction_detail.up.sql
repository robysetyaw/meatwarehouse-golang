CREATE TABLE transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INTEGER,
    meat_id INTEGER,
    meat_name TEXT,
    qty NUMERIC,
    price NUMERIC,
    total NUMERIC,
    is_active BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    created_by TEXT,
    updated_by TEXT
);
