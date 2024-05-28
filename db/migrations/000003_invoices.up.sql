CREATE TABLE IF NOT EXISTS invoices (
    game_id serial PRIMARY KEY,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    price Money
);
