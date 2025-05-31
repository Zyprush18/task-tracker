package test

import (
	"errors"

)

type TaskService struct {
	Repo TaskMock
}

func (service TaskService) Taskbyfilter(filter string) ([]Task, error) {
	task, err := service.Repo.Taskbyfilter(filter)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return task, nil
}


func (service TaskService) AddTask(task Task) (string, error) {
	add, err := service.Repo.AddTask(task)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return add, nil
}

func (service TaskService)	UpdateTask(id int,description string) (*Task,string, error) {
	task, err := service.Repo.FindById(id)
	if err != nil {
		return nil,"", errors.New(err.Error())
	}

	task.Description = description

	update, err := service.Repo.Update(task)
	if err != nil {
		return nil,"", errors.New(err.Error())
	}


	return &task,update,nil
}
