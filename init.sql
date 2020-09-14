
CREATE TABLE product (
    id          SERIAL          PRIMARY KEY,
    name        VARCHAR(15)     NOT NULL,
    type        VARCHAR(3)      NOT NULL,
    description VARCHAR(50)     NOT NULL,
    price       NUMERIC(3,2)   NOT NULL
);
