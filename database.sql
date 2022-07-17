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

create table if not exists contents
(
    id         BIGSERIAL primary key,
    data text,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table if not exists accounts
(
    id SERIAL primary key,
    name text,
    base_url text,
    api_key text,
    api_secret text,
    api_pass_phrase text,
    api_key_version int,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table if not exists strategy
(
    id SERIAL primary key,
    name text,
    data text,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table if not exists positions
(
    id SERIAL primary key,
    market_price float,
    side text,
    leverage text,
    position_size text,
    position_type text,
    symbol text,
    is_close bool,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);