BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id uuid primary key,
    created_at timestamp without time zone not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    full_name varchar(255) not null,
    show_count int,
    secret varchar(1500)
);

COMMIT;