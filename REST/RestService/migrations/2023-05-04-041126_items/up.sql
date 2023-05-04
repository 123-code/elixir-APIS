-- Your SQL goes here
CREATE TABLE Item(
    id serial NOT NULL,
    name varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    date_created timestamp NOT NULL,
    value int NOT NULL,
    CONSTRAINT items_pkey PRIMARY KEY (id)
);