-- +goose Up
-- +goose StatementBegin
create table accessible_role (
    id serial primary key,
    endpoint TEXT NOT NULL UNIQUE,
    role TEXT NOT NULL UNIQUE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table accessible_role;
-- +goose StatementEnd
