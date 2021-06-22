CREATE TABLE IF NOT EXISTS advert
(
    id              serial          not null unique,
    title           varchar(200)    not null,
    description     varchar(1000)   not null,
    price           numeric(12, 2)   not null,
    created_at      timestamp with time zone default current_timestamp
);

CREATE TABLE IF NOT EXISTS photo
(
    id              serial          not null unique,
    id_advert       int references advert (id) on delete cascade not null,
    url             varchar(200)    not null
    is_general      boolean         not null default false
);
