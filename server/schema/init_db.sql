CREATE TABLE IF NOT EXISTS urls
(
    short_url varchar(255) PRIMARY KEY,
    long_url varchar(255) not null unique
);