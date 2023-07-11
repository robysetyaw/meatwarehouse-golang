CREATE TABLE companies (
    id VARCHAR PRIMARY KEY,
    name VARCHAR,
    address VARCHAR,
    email VARCHAR,
    phone_number VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    created_by VARCHAR,
    updated_by VARCHAR
);