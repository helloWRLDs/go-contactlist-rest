CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    middle_name VARCHAR(30) NOT NULL,
    phone text
);

CREATE TABLE contact_in_group (
    id SERIAL PRIMARY KEY,
    contact_id INT,
    group_id INT,
    CONSTRAINT fk_group_id FOREIGN KEY (group_id) REFERENCES groups(id),
    CONSTRAINT fk_contact_id FOREIGN KEY (contact_id) REFERENCES contacts(id)
);

INSERT INTO groups(name) VALUES ('none');