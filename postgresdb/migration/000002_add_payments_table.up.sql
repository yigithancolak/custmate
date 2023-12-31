CREATE TYPE payment_methods AS ENUM ('credit_card', 'cash');
CREATE TYPE currencies AS ENUM ('try', 'usd', 'eur');



CREATE TABLE payments (
    id UUID PRIMARY KEY,
    amount INTEGER NOT NULL,
    payment_type payment_methods NOT NULL DEFAULT 'cash',
    currency currencies NOT NULL DEFAULT 'try',
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    customer_id UUID NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    org_group_id UUID NOT NULL REFERENCES org_groups(id) ON DELETE CASCADE,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE
);

CREATE INDEX idx_payments_date ON payments(date);
