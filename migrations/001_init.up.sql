CREATE TABLE IF NOT EXISTS users (
  id            uuid          NOT NULL, 
  email         varchar(255)  NOT NULL UNIQUE,
  login         varchar(80)   NOT NULL UNIQUE,
  first_name    varchar(20)   NOT NULL DEFAULT '',
  last_name     varchar(20)   NOT NULL DEFAULT '',
  middle_name   varchar(20)   NOT NULL DEFAULT '',
  avatar        varchar(255)  NOT NULL,
  created_at    timestamptz   NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  updated_at    timestamptz   NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  deleted_at    timestamptz   DEFAULT NULL, 
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS verifications (
  email             varchar(255)  NOT NULL,
  type              varchar(25)   NOT NULL,
  code              varchar(10)   NOT NULL,
  created_at        timestamptz   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at        timestamptz   NOT NULL,
  PRIMARY KEY (email),
  FOREIGN KEY (email) REFERENCES users (email) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS devices (
  id            uuid          NOT NULL,
  user_id       uuid          NOT NULL,
  device_token  varchar(256)  NOT NULL,
  created_at    timestamptz   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    timestamptz   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS access (
  id          uuid          NOT NULL,
  user_id     uuid          NOT NULL,
  device_id   uuid          NOT NULL,
  token       varchar(256)  NOT NULL,
  created_at  timestamptz   NOT NULL,
  expires_at  timestamptz   NOT NULL,
  deleted_at  timestamptz   DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id)   REFERENCES users (id) ON DELETE RESTRICT,
  FOREIGN KEY (device_id) REFERENCES devices (id) ON DELETE RESTRICT
);

CREATE SEQUENCE IF NOT EXISTS app_id_seq START 1;

CREATE TABLE IF NOT EXISTS apps (
  id     bigint       NOT NULL DEFAULT nextval('app_id_seq'),
  name   varchar(256) NOT NULL,
  secret varchar(256) NOT NULL,
  scope  varchar(25)  NOT NULL DEFAULT 'userinfo',   
  PRIMARY KEY (id)
);
