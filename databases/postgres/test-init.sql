INSERT INTO School(id, name) VALUES (0, 'Alfea'), (1, 'Hogwarts'), (2, 'New School'), (3,'Old School');

INSERT INTO Job(id, name) VALUES (0, 'Programmer'), (1, 'Cybersportsman'),
    (2, 'Gym instructor'), (3, 'Gangster');

INSERT INTO Person(id, name, school_id) VALUES (0, 'Bob', 3), (1, 'Harry', 1), (2, 'Bloom', 0),
    (3, 'Bib', 2), (4, 'Hardbass', 3), (5, 'Justin', 2), (6, 'Ivan', 1);

INSERT INTO JobLink(person_id, job_id) VALUES (0, 1), (0, 2), (0, 3), (1, 3), (2, 0), (2, 1),
    (2, 2), (2, 3), (3, 2), (3, 3), (4, 0), (4, 1), (4, 3), (5, 0), (5, 2), (5, 3);