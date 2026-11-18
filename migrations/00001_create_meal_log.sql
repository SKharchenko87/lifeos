-- +goose Up
create table meal_logs (
    id uuid primary key,
    created_at timestamptz not null default now(),
    user_id uuid not null,
    text text not null,
    kcal int not null,
    fat numeric(6, 2) not null,
    protein numeric(6, 2) not null,
    carbohydrate numeric(6, 2) not null,
    items jsonb not null default '[]'::jsonb
);
-- +goose Down
drop table if EXISTS meal_logs;