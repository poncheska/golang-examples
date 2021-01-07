package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
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

func main() {
	s, err := NewStorage("user=admin password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile("./test-init.sql")
	if err != nil {
		panic(err)
	}
	s.DB.MustExec(string(b))

	people := s.ReadPeople()
	jobs := s.ReadJobs()
	schools := s.ReadSchools()
	jobsNum := s.ReadJobsNumber()
	fmt.Printf("%v\n%v\n%v\n%v\n", people,
		jobs, schools, jobsNum)
	s.DB.Exec("INSERT INTO Person(id, name, school_id) VALUES (7, 'Fedor', 2);")
	people = s.ReadPeople()
	jobsNum = s.ReadJobsNumber()
	fmt.Printf("%v\n%v\n", people, jobsNum)
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
	s.DB.Select(&people, "SELECT * FROM Person;")
	return people
}

func (s *Storage) ReadSchools() []School {
	var schools []School
	s.DB.Select(&schools, "SELECT * FROM School;")
	return schools
}

func (s *Storage) ReadJobs() []Job {
	var jobs []Job
	s.DB.Select(&jobs, "SELECT * FROM Job;")
	return jobs
}

func (s *Storage) ReadJobsNumber() []JobsNumber {
	var jobsNum []JobsNumber
	s.DB.Select(&jobsNum, "SELECT * FROM JobsNumber;")
	return jobsNum
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
