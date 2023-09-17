package grades

import (
	"fmt"
	"sync"
)

//需要干什么（类型定义，全局变量，基本方法）

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

type GradeType string

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

type Students []Student

var (
	students      Students
	studentsMutex sync.Mutex
)

func (ss Students) GetByID(id int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}
	return nil, fmt.Errorf("Student with ID %d not found", id)
}

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
