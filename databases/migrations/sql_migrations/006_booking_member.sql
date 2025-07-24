-- +migrate Up
-- +migrate StatementBegin
create table adk_booking_member(
	id serial primary key,
	booking_id int not null,
	nama_penghuni varchar(128) not null,
	jenis_kelamin varchar(64) not null,
	nomor_handphone int not null,
	status_perkawinan varchar(64) not null,
	ktp_penghuni varchar(64) not null,
    created_at timestamp not null,
	created_by varchar(64) not null,
	modified_at timestamp,
	modified_by varchar(64),
    foreign key (booking_id) references adk_booking(id)
)
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
-- DROP TABLE IF EXISTS gorp_migrations;
-- DROP TABLE IF EXISTS users;
-- +migrate StatementEnd
