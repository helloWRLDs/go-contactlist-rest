CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    middle_name VARCHAR(30) NOT NULL,
    phone text,
    group_id int,
    CONSTRAINT fk_group_id FOREIGN KEY (group_id) REFERENCES groups(id)
);

INSERT INTO groups(name) VALUES ('none');