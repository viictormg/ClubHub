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
    owner_id INT REFERENCES owner(id),
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
    domain_contact_email VARCHAR(100),
    redirections JSON,
    number_redirections INT
);

INSERT INTO country (name, iso_code) VALUES
    ('Colombia', 'CO'),
    ('Mexito', 'MX'),
    ('Salvador', 'SV');

INSERT INTO owner (first_name, last_name, email, phone, city, country_code_id, address, zip_code) VALUES
    ('John', 'Doe', 'john.doe@example.com', '123456789', 'City1', 1, 'Address1', '12345'),
    ('Jane', 'Smith', 'jane.smith@example.com', '987654321', 'City2', 2, 'Address2', '54321'),
    ('Bob', 'Johnson', 'bob.johnson@example.com', '555555555', 'City3', 3, 'Address3', '67890');

INSERT INTO company (name, tax_number, city, country_code_id, owner_id, address, zip_code) VALUES
    ('Company1', '123456', 'City1', 1, 1, 'Address1', '12345'),
    ('Company2', '789012', 'City2', 2, 2, 'Address2', '54321'),
    ('Company3', '345678', 'City3', 3, 3, 'Address3', '67890');

INSERT INTO franchise (name, url, city, country_code_id, address, zip_code, company_id, protocol, logo_url, domain_creation, domain_expiration, domain_owner_name, domain_contact_email) VALUES
    ('Franchise1', 'https://www.franchise1.com', 'City1', 1, 'Address1', '12345', 1, 'HTTP', 'https://www.franchise1.com/logo.jpg', '2023-01-01', '2023-12-31', 'Owner1', 'owner1@example.com'),
    ('Franchise2', 'https://www.franchise2.com', 'City2', 2, 'Address2', '54321', 2, 'HTTPS', 'https://www.franchise2.com/logo.jpg', '2023-02-01', '2023-11-30', 'Owner2', 'owner2@example.com'),
    ('Franchise3', 'https://www.franchise3.com', 'City3', 3, 'Address3', '67890', 3, 'HTTP', 'https://www.franchise3.com/logo.jpg', '2023-03-01', '2023-10-31', 'Owner3', 'owner3@example.com');
