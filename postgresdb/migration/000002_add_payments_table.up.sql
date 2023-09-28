CREATE TYPE payment_methods AS ENUM ('credit_card', 'cash');
CREATE TYPE currencies AS ENUM ('try', 'usd', 'eur');



CREATE TABLE payments (
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL REFERENCES customers(id),
    amount INTEGER NOT NULL,
    payment_type payment_methods NOT NULL DEFAULT 'cash',
    currency currencies NOT NULL DEFAULT 'try',
    date DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE INDEX idx_payments_date ON payments(date);
