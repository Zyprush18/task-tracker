package test

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// entity
type Task struct {
	Id string
	Description string 
	Status string
	CreatedAt time.Time	
	UpdatedAt time.Time
}


type TaskMock interface {
	ListAllTask() ([]Task, error)
}

// mock object
type TaskMockObject struct {
	Mock mock.Mock
}

func (m *TaskMockObject) ListAllTask() ([]Task, error) {
	args := m.Mock.Called()
	return args.Get(0).([]Task), args.Error(1)
}