package db

const CreateUsersAccount = `Create table  if not exists users(
	id bigserial primary key,
	name text not null,
	surname text not null,
	age integer not null,
	gender text not null,
	login text not null unique,
	password text not null,
	remove boolean not null default false
);`

const CreateATMsTable = `Create table  if not exists atms(
	id bigserial primary key,
	name text not null,
	status boolean not null default true
);`