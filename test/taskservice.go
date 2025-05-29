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