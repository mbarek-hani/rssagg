-- +goose Up

create table users (
    id uuid not null primary key,
    name text not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

-- +goose Down

drop table users;
