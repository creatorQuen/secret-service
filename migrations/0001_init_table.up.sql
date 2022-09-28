BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL primary key,
    created_at timestamp with time zone not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    full_name varchar(255) not null,
    show_count int,
    unique_id varchar(255),
    secret varchar(1500)
);

COMMIT;