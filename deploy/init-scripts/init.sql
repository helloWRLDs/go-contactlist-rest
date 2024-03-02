CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    modified_at TIMESTAMP
);

CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    phone text NOT NULL,
    created_at TIMESTAMP,
    modified_at TIMESTAMP
);

CREATE TABLE contact_in_group (
    id SERIAL PRIMARY KEY,
    contact_id INT,
    group_id INT,
    CONSTRAINT fk_group_id FOREIGN KEY (group_id) REFERENCES groups(id),
    CONSTRAINT fk_contact_id FOREIGN KEY (contact_id) REFERENCES contacts(id)
);