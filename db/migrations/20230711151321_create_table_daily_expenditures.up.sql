CREATE TABLE daily_expenditures (
    id VARCHAR PRIMARY KEY,
    user_id VARCHAT,
    amount NUMERIC,
    description VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    created_by VARCHAR,
    updated_by VARCHAR
);