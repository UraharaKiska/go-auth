-- +goose Up
-- +goose StatementBegin
create table "users" (
    id serial primary key,
    name text not null UNIQUE,
    email text not null UNIQUE,
    password_hash text not null,
    role text not null,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp,
    is_active BOOLEAN DEFAULT TRUE 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "users";
-- +goose StatementEnd
