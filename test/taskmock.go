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
	Taskbyfilter(filter string) ([]Task, error)
	AddTask(task Task) (string, error)
	FindById(id int) (Task, error)
	Update(task Task) (string, error)
	Delete(task Task) (string, error)
	MarkStatus(task Task) (string, error)
}
// mock object
type TaskMockObject struct {
	Mock mock.Mock
}

func (m *TaskMockObject) ListAllTask() ([]Task, error) {
	args := m.Mock.Called()
	return args.Get(0).([]Task), args.Error(1)
}


func (m *TaskMockObject) Taskbyfilter(filter string) ([]Task, error) {
	args := m.Mock.Called(filter)

	return args.Get(0).([]Task), args.Error(1)
}

func (m *TaskMockObject) AddTask(task Task) (string, error) {
	args := m.Mock.Called(task)
	return args.String(0), args.Error(1)
}

func (m *TaskMockObject) FindById(id int) (Task, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(Task), args.Error(1)
}

func (m *TaskMockObject) Update(task Task) (string, error) {
	args := m.Mock.Called(task)

	return args.String(0),args.Error(1)
}


func (m *TaskMockObject) Delete(task Task) (string, error) {
	args := m.Mock.Called(task)

	return args.String(0), args.Error(1)
}

func (m *TaskMockObject) MarkStatus(task Task) (string, error) {
	args := m.Mock.Called(task)

	return args.String(0), args.Error(1)
}