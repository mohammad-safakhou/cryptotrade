create table if not exists content
(
    id SERIAL primary key,
    request_id int,
    content text,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);