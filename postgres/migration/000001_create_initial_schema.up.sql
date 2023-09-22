CREATE TABLE organizations {
    id UUID PRIMARY KEY
    name VARCHAR NOT NULL
    email VARCHAR NOT NULL
    password VARCHAR NOT NULL
}

CREATE TABLE instructors {
    id UUID PRIMARY KEY
    name VARCHAR NOT NULL
    organization_id UUID NOT NULL REFERENCES organizations(id)
}

CREATE TABLE times {
    id UUID PRIMARY KEY
    group_id UUID NOT NULL REFERENCES groups(id) --suspicious TODO: decide
    day VARCHAR NOT NULL
    start_hour VARCHAR NOT NULL
    finish_hour VARCHAR NOT NULL
}


CREATE TABLE groups {
    id UUID PRIMARY KEY
    name VARCHAR NOT NULL
    organization_id UUID NOT NULL REFERENCES organizations(id)
    instructor_id UUID NOT NULL REFERENCES instructors(id)
    times_id UUID NOT NULL REFERENCES times(id)
}

CREATE TABLE customers {
    id UUID PRIMARY KEY
    name  VARCHAR NOT NULL
    organization_id UUID NOT NULL REFERENCES organizations(id)


}

--go on from there