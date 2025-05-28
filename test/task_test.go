package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

)

func TestListAllTask(t *testing.T)  {
	m := new(TaskMockObject)
	m.Mock.On("ListAllTask").Return([]Task{
		{
			Id: "1",
			Description: "Description 1",
			Status: "todo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id: "2",
			Description: "Description 2",
			Status: "in-progres",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id: "3",
			Description: "Description 3",
			Status: "done",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	},nil)
	
	task, err := m.ListAllTask()
	assert.NoError(t, err)
	assert.Len(t,task, 3)

	assert.Equal(t, "1", task[0].Id)

	m.Mock.AssertExpectations(t)
}