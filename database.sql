create table if not exists candles_content
(
    id         BIGSERIAL primary key,
    time_frame bigint,
    opening    float8,
    closing    float8,
    highest    float8,
    lowest     float8,
    volume     float8,
    amount     float8,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
