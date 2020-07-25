CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY NOT NULL,
    name varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);