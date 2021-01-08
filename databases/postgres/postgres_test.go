package postgres

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func init(){
	s, err := NewStorage("user=admin password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile("./test-init.sql")
	if err != nil {
		panic(err)
	}
	s.DB.MustExec(string(b))
}

func TestStorage_GetSchName(t *testing.T){
	s, err := NewStorage("user=admin password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	schnames := []string{"Old School","Hogwarts","Alfea",
		"New School","Old School","New School","Hogwarts"}
	for i, v := range schnames{
		n, err := s.GetSchName(int64(i))
		assert.Nil(t, err)
		assert.Equal(t, v, n)
	}

	_, err = s.GetSchName(int64(10))
	assert.NotNil(t, err)
}

func TestStorage_ReadJobs(t *testing.T) {
	s, err := NewStorage("user=admin password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	jobs := s.ReadJobs()
	assert.Equal(t, 4, len(jobs))
}

func TestStorage_ReadPeople(t *testing.T) {
	s, err := NewStorage("user=admin password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	people := s.ReadPeople()
	assert.Equal(t, 7, len(people))
}

func TestStorage_ReadSchools(t *testing.T) {
	s, err := NewStorage("user=admin password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	schools := s.ReadSchools()
	assert.Equal(t, 4, len(schools))
}

func TestStorage_ReadJobsNumber(t *testing.T) {
	s, err := NewStorage("user=admin password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	jn := s.ReadJobsNumber()
	var sum int64
	for _, v := range jn{
		sum += v.Id
	}
	assert.Equal(t, int64(16), sum)
}