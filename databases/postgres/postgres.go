package postgres

import "github.com/jmoiron/sqlx"

type Storage struct {
	DB sqlx.DB
}

type Person struct {
	Id       int64 `db:"id"`
	Name     string `db:"name"`
	SchoolId int64 `db:"school_id"`
}

type Job struct {
	Id   int64 `db:"id"`
	Name string `db:"name"`
}

type School struct {
	Id   int64 `db:"id"`
	Name string `db:"name"`
}

type JobsNumber struct {
	Id   int64 `db:"id"`
	Name string `db:"name"`
	Num  int64 `db:"num"`
}
