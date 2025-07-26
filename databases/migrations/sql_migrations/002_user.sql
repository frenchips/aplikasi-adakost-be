-- +migrate Up
-- +migrate StatementBegin
create table adk_users(
	id serial primary key,
	username varchar(64) not null,
	password TEXT NOT NULL,
    fullname VARCHAR(100),
    no_handphone VARCHAR(20),
    email VARCHAR(100),
    role_id int not null,
	created_at timestamp not null,
	created_by varchar(64) not null,
	modified_at timestamp,
	modified_by varchar(64),
    foreign key (role_id) references adk_roles(id)
)
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
-- DROP TABLE IF EXISTS gorp_migrations;
-- DROP TABLE IF EXISTS users;
-- +migrate StatementEnd
