-- +goose Up
-- +goose StatementBegin
create table auth (
    id serial primary key,
    name text not null,
    email text not null,
    password text not null,
    password_confirm text not null,
    role text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table auth;
-- +goose StatementEnd
