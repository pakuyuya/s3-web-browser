CREATE extension pgcrypto;

CREATE SCHEMA s3web;

CREATE TABLE s3web.users(
    id SERIAL NOT NULL,
    username VARCHAR(128),
    loginid VARCHAR(60),
    password_sha256 VARCHAR(128),
    permissionsjson VARCHAR(512),
    create_at TIMESTAMP,
    update_at TIMESTAMP
);


ALTER TABLE s3web.users
  ADD CONSTRAINT s3web_users_pkey PRIMARY KEY (id);

CREATE INDEX s3web_users_idx_1 ON s3web.users (loginid);

INSERT INTO s3web.users (username, loginid, password_sha256, permissionsjson, create_at, update_at)
  VALUES ('Admin', 'admin', digest('password', 'sha256'), '{"admin":true}', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

CREATE TABLE s3web.profiles(
    profileid SERIAL NOT NULL,
    profilename VARCHAR(32),
    connjson VARCHAR(1024),
    bucket VARCHAR(128),
    basepath VARCHAR(512),
    create_at TIMESTAMP,
    update_at TIMESTAMP
);

ALTER TABLE s3web.profiles
  ADD CONSTRAINT s3web_profiles_pkey PRIMARY KEY (profileid);
