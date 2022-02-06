CREATE DATABASE shortener WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';


\connect shortener

CREATE TABLE IF NOT EXISTS links
(
    id serial PRIMARY KEY,
    link text NOT NULL,
    token VARCHAR (25)
);

CREATE TABLE IF NOT EXISTS links_stat
(
    id         serial PRIMARY KEY,
    link_id    int     NOT NULL constraint links_stat_links_id_fk references links,
    ip         varchar NOT NULL,
    created_at timestamp default current_timestamp
);

CREATE ROLE short WITH LOGIN PASSWORD 'iniT11';
GRANT ALL PRIVILEGES ON DATABASE shortener TO short;
GRANT ALL ON links to short;
GRANT ALL ON links_id_seq to short;