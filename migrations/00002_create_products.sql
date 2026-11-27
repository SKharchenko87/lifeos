-- +goose Up
CREATE EXTENSION IF NOT EXISTS pg_trgm;

create table products (
    id uuid primary key default gen_random_uuid(),
    name text not null,
    brand text,
    category text,
    barcode varchar(32),
    kcal int not null,
    fat numeric(6,2) not null,
    protein numeric(6,2) not null,
    carbohydrate numeric(6,2) not null,
    source text not null default 'manual',
    created_at timestamptz not null default now()
);

create index idx_products_name_trgm on products using gin (name gin_trgm_ops);


-- +goose Down

drop index if exists idx_products_name_trgm;

drop table if EXISTS products;
