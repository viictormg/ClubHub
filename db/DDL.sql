CREATE TABLE country (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    iso_code CHAR(2) NOT NULL
);

CREATE TABLE owner (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    city VARCHAR(100) NOT NULL,
    country_code_id INT REFERENCES country(id),
    address VARCHAR(255) NOT NULL,
    zip_code VARCHAR(20) NOT NULL
);


CREATE TABLE company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    tax_number VARCHAR(20) NOT NULL,
    city VARCHAR(100) NOT NULL,
    country_code_id INT REFERENCES country(id),
    owner_id INT REFERENCES owner(owner_id),
    address VARCHAR(255) NOT NULL,
    zip_code VARCHAR(20) NOT NULL
);


CREATE TABLE franchise (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    url VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    country_code_id INT REFERENCES country(id),
    address VARCHAR(255) NOT NULL,
    zip_code VARCHAR(20) NOT NULL,
    company_id INT REFERENCES company(id),
    protocol VARCHAR(10) NOT NULL,
    logo_url VARCHAR(255),
    domain_creation DATE,
    domain_expiration DATE,
    domain_owner_name VARCHAR(100),
    domain_contact_email VARCHAR(100)
);

