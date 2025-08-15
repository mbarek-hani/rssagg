-- +goose Up

create table feeds (
    id uuid not null primary key,
    name text not null,
    url text not null unique,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id uuid references users(id) on delete cascade
);

-- +goose Down

drop table feeds;
