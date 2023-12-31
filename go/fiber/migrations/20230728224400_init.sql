-- DDL generated by Postico 2.0.4
-- Not all database features are supported. Do not use for backup.

-- Table Definition ----------------------------------------------

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    activation_code character varying(36),
    activation_code_expires_at timestamp(0) without time zone DEFAULT (now() + '3 days'::interval),
    inserted_at timestamp(0) without time zone NOT NULL,
    updated_at timestamp(0) without time zone NOT NULL
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX users_pkey ON users(id int8_ops);
CREATE UNIQUE INDEX users_email_index ON users(email text_ops);
CREATE UNIQUE INDEX users_activation_code_index ON users(activation_code text_ops);
