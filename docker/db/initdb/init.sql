CREATE SCHEMA s3web;

CREATE TABLE s3web.users(
    id SERIAL NOT NULL,
    display VARCHAR(128),
    loginid VARCHAR(60),
    password_sha256 VARCHAR(64),
    create_at TIMESTAMP,
    update_at TIMESTAMP
);

INSERT INTO s3web.users (display, loginid, password_sha256, create_at, update_at)
  VALUES ('Admin', 'admin', digest('password', 'sha256'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

CREATE TABLE s3web.profiles(
    profileid VARCHAR(32),
    profilename VARCHAR(128),
    connjson VARCHAR(1024),
    bucket VARCHAR(128),
    basepath VARCHAR(523),
    create_at TIMESTAMP,
    update_at TIMESTAMP
);
