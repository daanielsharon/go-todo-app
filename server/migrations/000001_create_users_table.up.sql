CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(30) UNIQUE NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
);
