CREATE extension pgcrypto;

CREATE SCHEMA s3web;

CREATE TABLE s3web.users(
    id SERIAL NOT NULL,
    username VARCHAR(128),
    loginid VARCHAR(60),
    password_sha256 VARCHAR(128),
    create_at TIMESTAMP,
    update_at TIMESTAMP
);

INSERT INTO s3web.users (username, loginid, password_sha256, create_at, update_at)
  VALUES ('Admin', 'admin', digest('password', 'sha256'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

CREATE TABLE s3web.profiles(
    profileid VARCHAR(32),
    profilename VARCHAR(32),
    connjson VARCHAR(1024),
    bucket VARCHAR(128),
    basepath VARCHAR(512),
    create_at TIMESTAMP,
    update_at TIMESTAMP
);
