#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    description VARCHAR,
    price_in_cents INT
);
INSERT INTO products VALUES (1, 'Super NES Classic', 'The Super NES Classic Edition', 9149);
INSERT INTO products VALUES (2, 'Marvel’s Spider-Man', 'PlayStation 4 game', 5499);
INSERT INTO products VALUES (3, 'Nintendo Switch', 'Neon Red and Neon Blue Joy-Con', 299);

CREATE TABLE customers (id SERIAL PRIMARY KEY, first_name VARCHAR, last_name VARCHAR, birth_date DATE);
INSERT INTO customers VALUES (1, 'Myrcella', 'Baratheon', '1999-10-01');
INSERT INTO customers VALUES (2, 'Daenerys', 'Targaryen', CURRENT_DATE);
INSERT INTO customers VALUES (3, 'Margaery', 'Tyrell', '1982-02-11');
INSERT INTO customers VALUES (4, 'Ellaria', 'Sand', '1973-05-14');

CREATE TABLE discounts (product_id INT, max_discount DECIMAL);
INSERT INTO discounts VALUES (1, 0.1);
INSERT INTO discounts VALUES (2, 0.05);
EOSQL