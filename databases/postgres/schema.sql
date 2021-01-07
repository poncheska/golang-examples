CREATE TABLE School(
                       id serial NOT NULL PRIMARY KEY,
                       name varchar(30) NOT NULL UNIQUE
);

CREATE TABLE Person(
    id serial NOT NULL PRIMARY KEY,
    name varchar(30) NOT NULL,
    school_id integer REFERENCES school(id)
);

CREATE TABLE Job(
    id serial NOT NULL PRIMARY KEY,
    name varchar(30) NOT NULL UNIQUE
);

CREATE TABLE JobLink(
    person_id integer REFERENCES Person(id),
    job_id integer REFERENCES Job(id)
);

CREATE VIEW JobsNumber AS
    SELECT P.id AS id, P.name AS name, COUNT(JL.job_id) AS num
    FROM Person P
    JOIN JobLink JL on P.id = JL.person_id
    GROUP BY (P.id, P.name)
    UNION
    SELECT P.id AS id, P.name AS name, 0 AS num
    FROM Person P
    WHERE P.id NOT IN (
        SELECT person_id FROM JobLink
        );