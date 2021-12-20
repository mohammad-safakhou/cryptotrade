create table if not exists candles_content
(
    id SERIAL primary key,
    time_frame    int,
	opening double,
	closing double,
	highest double,
	lowest  double,
	volume  double,
	amount  double,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
