-- +goose Up
-- +goose StatementBegin
create table access_blacklist (
    id serial primary key,
    token TEXT NOT NULL UNIQUE,
    created_at timestamp not null default CURRENT_TIMESTAMP
);

create table refresh_blacklist (
    id serial primary key,
    token TEXT NOT NULL UNIQUE,
    created_at timestamp not null default CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table access_blacklist;
drop table refresh_blacklist;
-- +goose StatementEnd
