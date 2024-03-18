CREATE DATABASE service_users;
\c service_users;
CREATE TABLE users (
    id VARCHAR(225) PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(225) NOT NULL
);

CREATE DATABASE service_employees;
\c service_employees;

CREATE TABLE employees (
    id VARCHAR(225) PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);