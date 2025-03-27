CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    justification VARCHAR(255) NOT NULL,
    type VARCHAR(50) CHECK (type IN ('investment', 'income', 'outcome')) NOT NULL,
    category VARCHAR(100) NOT NULL,
    start_date BIGINT NOT NULL,
    end_date BIGINT NOT NULL,
    frequency VARCHAR(50) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL
);