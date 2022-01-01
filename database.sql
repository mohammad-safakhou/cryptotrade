create table if not exists candles_content
(
    id         BIGSERIAL primary key,
    time_frame float8,
    opening    text,
    closing    text,
    highest    text,
    lowest     text,
    volume     text,
    amount     text,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
