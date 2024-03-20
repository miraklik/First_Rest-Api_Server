CREATE TABLE users (
    id BIGINT not null primary key,
    email VARCHAR not null unique,
    encrypto_password VARCHAR not null
);

