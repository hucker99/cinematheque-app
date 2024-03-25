DROP TABLE IF EXISTS actors;
CREATE TABLE actors
(
    id        serial          NOT NULL UNIQUE,
    name      varchar(255)    NOT NULL,
    gender    varchar(10)     NOT NULL,
    birthday  varchar(255)    NOT NULL
);

DROP TABLE IF EXISTS films;
CREATE TABLE films
(
    id            serial          NOT NULL UNIQUE,
    title         varchar(255)    NOT NULL,
    release_date  varchar(255)    NOT NULL,
    rating        varchar(255)    NOT NULL
);

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id              serial          NOT NULL UNIQUE,
    email           varchar(255)    NOT NULL,
    password_hash   varchar(255)    NOT NULL,
    role            varchar(10)     NOT NULL
);

DROP TABLE IF EXISTS films_actors;
CREATE TABLE films_actors
(
    id          serial                                          NOT NULL UNIQUE,
    film_id     int REFERENCES films (id) ON DELETE CASCADE     NOT NULL,
    actor_id    int REFERENCES actors (id) ON DELETE CASCADE    NOT NULL
);