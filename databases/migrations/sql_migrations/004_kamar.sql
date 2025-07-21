-- +migrate Up
-- +migrate StatementBegin
create table adk_kamar(
	id serial primary key,
	kost_id int not null,
	nomor_kamar varchar(64) not null,
    harga_per_bulan int not null,
    status_kamar varchar(64) not null,
    created_at timestamp not null,
	created_by varchar(64) not null,
	modified_at timestamp,
	modified_by varchar(64),
    foreign key (kost_id) references adk_kost(id)
)
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
-- DROP TABLE IF EXISTS gorp_migrations;
-- DROP TABLE IF EXISTS users;
-- +migrate StatementEnd
