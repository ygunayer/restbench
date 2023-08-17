CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  status varchar(32) NOT NULL default('pending_activation'),
  email varchar(255) NOT NULL,
  password_hash varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  activation_code varchar(36),
  activation_code_expires_at timestamp(0) without time zone DEFAULT (now() + '3 days'::interval),
  inserted_at timestamp(0) without time zone NOT NULL,
  updated_at timestamp(0) without time zone NOT NULL
);

CREATE UNIQUE INDEX users_email_index ON users(email);
CREATE UNIQUE INDEX users_activation_code_index ON users(activation_code);
