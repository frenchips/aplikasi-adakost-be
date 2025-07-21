-- +migrate Up
-- +migrate StatementBegin
create table adk_kost(
	id serial primary key,
	nama_kost varchar(64) not null,
	alamat text not null,
    pemilik_id int not null,
	type_kost varchar(64) not null,
    created_at timestamp not null,
	created_by varchar(64) not null,
	modified_at timestamp,
	modified_by varchar(64),
    foreign key (pemilik_id) references adk_users(id)
)
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
-- DROP TABLE IF EXISTS gorp_migrations;
-- DROP TABLE IF EXISTS users;
-- +migrate StatementEnd
