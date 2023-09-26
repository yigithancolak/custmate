CREATE TABLE organizations (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,  
    password VARCHAR NOT NULL  
);

CREATE TABLE instructors (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE
);

CREATE TABLE org_groups (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    instructor_id UUID NOT NULL REFERENCES instructors(id) ON DELETE CASCADE
);

CREATE TYPE days_of_week AS ENUM ('monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday', 'sunday');

CREATE TABLE times (
    id UUID PRIMARY KEY,
    org_group_id UUID NOT NULL REFERENCES org_groups(id) ON DELETE CASCADE,
    day days_of_week NOT NULL,
    start_hour TIME NOT NULL,
    finish_hour TIME NOT NULL
);

CREATE TABLE customers (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE customer_groups (
    customer_id UUID NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    org_group_id UUID NOT NULL REFERENCES org_groups(id) ON DELETE CASCADE,
    PRIMARY KEY (customer_id, org_group_id)
);

CREATE INDEX idx_customers_name ON customers(name);
