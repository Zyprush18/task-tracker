package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListAllTask(t *testing.T) {
	m := new(TaskMockObject)
	m.Mock.On("ListAllTask").Return([]Task{
		{
			Id:          "1",
			Description: "Description 1",
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          "2",
			Description: "Description 2",
			Status:      "in-progres",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          "3",
			Description: "Description 3",
			Status:      "done",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil)

	task, err := m.ListAllTask()
	assert.NoError(t, err)
	assert.Len(t, task, 3)

	assert.Equal(t, "1", task[0].Id)

	m.Mock.AssertExpectations(t)
}

var Taskmock = &TaskMockObject{Mock: mock.Mock{}}
var Taskservice = TaskService{Repo: Taskmock}

func TestListbyfilter(t *testing.T) {
	status := []string{"todo","in-progress","done"}

	data := []Task{
		{
			Id:          "1",
			Description: "Description 1",
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          "2",
			Description: "Description 2",
			Status:      "in-progress",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          "3",
			Description: "Description 3",
			Status:      "done",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          "4",
			Description: "Description 4",
			Status:      "done",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          "5",
			Description: "Description 5",
			Status:      "done",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for i, d := range data {
		for _, s := range status {
			if d.Status == s {
				t.Run(d.Status, func(t *testing.T) {
					Taskmock.Mock.On("Taskbyfilter", d.Status).Return(data, nil)
					result, err := Taskservice.Repo.Taskbyfilter(data[i].Status)
		
					assert.NoError(t, err)
					assert.Equal(t, d.Status, result[i].Status)
				})
			}
		}
	}


	Taskmock.Mock.AssertExpectations(t)
}
