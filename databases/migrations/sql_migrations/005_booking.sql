-- +migrate Up
-- +migrate StatementBegin
create table adk_booking(
	id serial primary key,
	user_id int not null,
	kamar_id int not null,
    tanggal_mulai timestamp not null,
    tanggal_akhir timestamp ,
    jumlah_penghuni int not null,
    status_booking varchar(64) not null,
    created_at timestamp not null,
	created_by varchar(64) not null,
	modified_at timestamp,
	modified_by varchar(64),
    foreign key (user_id) references adk_users(id),
    foreign key (kamar_id) references adk_kamar(id)
)
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
-- DROP TABLE IF EXISTS gorp_migrations;
-- DROP TABLE IF EXISTS users;
-- +migrate StatementEnd
