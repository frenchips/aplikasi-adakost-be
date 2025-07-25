DROP TABLE IF EXISTS gorp_migrations;

-- +migrate Up
-- +migrate StatementBegin
create table adk_roles(
	id serial primary key,
    name varchar(64) not null,
    created_at timestamp not null,
	created_by varchar(64) not null,
	modified_at timestamp,
	modified_by varchar(64)
);

INSERT INTO adk_roles (id, name, created_at, created_by)
VALUES
  (1, 'penyewa', now(), 'SYSTEM'),
  (2, 'pemilik', now(), 'SYSTEM')
 
-- +migrate StatementEnd

