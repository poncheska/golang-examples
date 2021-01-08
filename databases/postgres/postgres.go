package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sqlx.DB
}

type Person struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	SchoolId int64  `db:"school_id"`
}

type Job struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

type School struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

type JobsNumber struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Num  int64  `db:"num"`
}

func NewStorage(connStr string) (*Storage, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Storage{db}, nil
}

func (s *Storage) ReadPeople() []Person {
	var people []Person
	err := s.DB.Select(&people, "SELECT * FROM Person;")
	if err != nil {
		panic(err)
	}
	return people
}

func (s *Storage) ReadSchools() []School {
	var schools []School
	err := s.DB.Select(&schools, "SELECT * FROM School;")
	if err != nil {
		panic(err)
	}
	return schools
}

func (s *Storage) ReadJobs() []Job {
	var jobs []Job
	err := s.DB.Select(&jobs, "SELECT * FROM Job;")
	if err != nil {
		panic(err)
	}
	return jobs
}

func (s *Storage) ReadJobsNumber() []JobsNumber {
	var jobsNum []JobsNumber
	err := s.DB.Select(&jobsNum, "SELECT * FROM JobsNumber;")
	if err != nil {
		panic(err)
	}
	return jobsNum
}

func (s *Storage) GetSchName(personId int64) (string, error) {
	type SchName struct {
		Name string `db:"name"`
	}
	var sn SchName
	err := s.DB.Get(&sn,
		`SELECT S.name AS name FROM School S
				JOIN Person P ON P.school_id=S.id
				WHERE P.id=$1`, personId)
	if err != nil{
		return "", err
	}
	return sn.Name, nil
}

// This func doesn't work correct, because init.sql insert data with id
// and serial counter doesn't know about it.
//
//func (s *Storage) InsertPerson(name string, schoolId int64) (*Person, error){
//	res, err := s.DB.Exec("INSERT INTO Person(name, school_id) VALUES ($1, $2);", name, schoolId)
//	if err != nil {
//		return nil, err
//	}
//	id, err := res.LastInsertId()
//	if err != nil {
//		return nil, err
//	}
//	return &Person{
//		Id:       id,
//		Name:     name,
//		SchoolId: schoolId,
//	}, nil
//}
