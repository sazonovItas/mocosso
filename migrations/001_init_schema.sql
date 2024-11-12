-- +goose Up
CREATE SCHEMA IF NOT EXISTS sso;

SET SEARCH_PATH TO sso, PUBLIC;

CREATE TABLE IF NOT EXISTS user_account (
  id            bigserial       NOT NULL,
  email         varchar(255)    NOT NULL UNIQUE,
  username      varchar(70)     NOT NULL UNIQUE, 
  password_hash varchar(255)    NOT NULL,
  avatar        text            DEFAULT NULL,
  is_verified   pg_catalog.bool DEFAULT FALSE,
  created_at    timestamptz     NOT NULL DEFAULT NOW(),
  updated_at    timestamptz     NOT NULL DEFAULT NOW(),
  deleted_at    timestamptz     DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS verification (
  email       varchar(255)  NOT NULL,
  type        varchar(25)   NOT NULL,
  code        varchar(8)    NOT NULL,
  token       varchar(64)   NOT NULL UNIQUE,
  expires_at  timestamptz   NOT NULL,
  updated_at  timestamptz   NOT NULL DEFAULT NOW(),
  PRIMARY KEY (email),
  FOREIGN KEY (email) REFERENCES user_account (email)
);

CREATE TABLE IF NOT EXISTS device (
  id            bigserial     NOT NULL,
  name          text          NOT NULL,
  user_id       bigint        NOT NULL,
  hash_id       varchar(64)   NOT NULL UNIQUE,
  last_used_at  timestamptz   NOT NULL DEFAULT NOW(),
  created_at    timestamptz   NOT NULL DEFAULT NOW(),
  updated_at    timestamptz   NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES user_account (id)
);

CREATE TABLE IF NOT EXISTS access (
  id            bigserial   NOT NULL,
  user_id       bigint      NOT NULL,
  device_id     bigint      NOT NULL,
  refresh_token text        NOT NULL,
  created_at    timestamptz NOT NULL DEFAULT NOW(),
  last_used_at  timestamptz NOT NULL DEFAULT NOW(),
  expires_at    timestamptz NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (device_id) REFERENCES device (id),
  FOREIGN KEY (user_id) REFERENCES user_account (id)
);

CREATE TABLE IF NOT EXISTS role (
  id          serial      NOT NULL,
  name        varchar(70) NOT NULL UNIQUE,
  description text        DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS scope ( 
  id          serial      NOT NULL,
  name        varchar(70) NOT NULL UNIQUE,
  description text        DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_role (
  user_id bigint  NOT NULL,
  role_id int     NOT NULL,
  PRIMARY KEY (user_id, role_id),
  FOREIGN KEY (user_id) REFERENCES user_account (id),
  FOREIGN KEY (role_id) REFERENCES role (id)
);

CREATE TABLE IF NOT EXISTS role_scope (
  role_id  int NOT NULL,
  scope_id int NOT NULL,
  PRIMARY KEY (role_id, scope_id),
  FOREIGN KEY (role_id) REFERENCES role (id),
  FOREIGN KEY (scope_id) REFERENCES scope (id)
);

-- +goose Down
DROP SCHEMA IF EXISTS sso CASCADE;
