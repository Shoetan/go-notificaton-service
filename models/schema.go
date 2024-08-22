package models


var TABLES = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name text NOT NULL,
		last_name text NOT NULL,
		email_address text NOT NULL,
		password varchar,
		inserted_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at timestamp NULL
	)`
