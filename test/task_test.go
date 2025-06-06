package test

import (
	"errors"
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
	status := []string{"todo", "in-progress", "done"}

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

func TestAddTask(t *testing.T) {
	data := Task{
		Id:          "1",
		Description: "Ini Description 1",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	Taskmock.Mock.On("AddTask", data).Return("Task added successfully", nil)

	result, err := Taskservice.Repo.AddTask(data)

	assert.NoError(t, err)

	assert.Equal(t, "Task added successfully", result)

	Taskmock.Mock.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	data := Task{
		Id:          "1",
		Description: "Ini Description 1",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	t.Run("UpdateTask", func(t *testing.T) {
		Taskmock.Mock.On("FindById", 1).Return(data, nil)
		Taskmock.Mock.On("Update", mock.AnythingOfType("Task")).Return("Task Updated successfully", nil)

		result, message, err := Taskservice.UpdateTask(1, "Ini Description 1 Update")
		assert.NoError(t, err)

		assert.Equal(t, "Task Updated successfully", message)
		assert.Equal(t, "Ini Description 1 Update", result.Description)
	})

	t.Run("UpdateTaskNotFound", func(t *testing.T) {
		Taskmock.Mock.On("FindById", 2).Return(data, errors.New("Task with Id 2 Not Found"))
		Taskmock.Mock.On("Update", mock.AnythingOfType("Task")).Return("", errors.New("Failed To Update Task"))

		result, message, err := Taskservice.UpdateTask(2, "Ini Description 2 Update")
		assert.Error(t, err)

		assert.Equal(t, "Task with Id 2 Not Found", err.Error())
		assert.Equal(t, "", message)
		assert.Nil(t, result)
	})

	Taskmock.Mock.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	data := Task{
		Id:          "1",
		Description: "Ini Description 1",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("DeleteTask", func(t *testing.T) {
		Taskmock.Mock.On("FindById", 1).Return(data, nil)
		Taskmock.Mock.On("Delete", mock.AnythingOfType("Task")).Return("Task Deleted successfully", nil)

		result, err := Taskservice.DeleteTask(1)

		assert.NoError(t, err)
		assert.Equal(t, "Task Deleted successfully", result)
	})

	t.Run("DeleteTaskNotFound", func(t *testing.T) {
		Taskmock.Mock.On("FindById", 3).Return(data, errors.New("Task with Id 3 Not Found"))
		Taskmock.Mock.On("Delete", mock.AnythingOfType("Task")).Return("", errors.New("Failed To Delete Task"))

		msg, err := Taskservice.DeleteTask(3)

		assert.Error(t, err)
		assert.Equal(t, "Task with Id 3 Not Found", err.Error())
		assert.Equal(t, "", msg)
	})

	Taskmock.Mock.AssertExpectations(t)
}


func TestMarkProgress(t *testing.T)  {
	data := Task{
		Id:          "1",
		Description: "Ini Description 1",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("In-Progress",func(t *testing.T) {
		Taskmock.Mock.On("FindById", 1).Return(data, nil)
		Taskmock.Mock.On("MarkStatus", mock.AnythingOfType("Task")).Return("Task Updated successfully",nil)

		msg, err := Taskservice.MarkUpdate(1,"in-progress")

		assert.NoError(t, err)
		assert.Equal(t, "Task Updated successfully", msg)
	})


	t.Run("Done",func(t *testing.T) {
		Taskmock.Mock.On("FindById", 1).Return(data, nil)
		Taskmock.Mock.On("MarkStatus", mock.AnythingOfType("Task")).Return("Task Updated successfully",nil)

		msg, err := Taskservice.MarkUpdate(1,"done")

		assert.NoError(t, err)
		assert.Equal(t, "Task Updated successfully", msg)
	})

	t.Run("NotFoundDataById",func(t *testing.T) {
		Taskmock.Mock.On("FindById", 2).Return(data, errors.New("Task with Id 2 Not Found"))
		Taskmock.Mock.On("MarkStatus", mock.AnythingOfType("Task")).Return("",errors.New("Failed To Update Task"))

		msg, err := Taskservice.MarkUpdate(2,"done")

		assert.Error(t, err)
		assert.Equal(t, "", msg)
		assert.Equal(t, "Task with Id 2 Not Found", err.Error())
	})

	Taskmock.Mock.AssertExpectations(t)
}