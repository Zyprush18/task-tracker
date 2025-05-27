package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Zyprush18/task-tracker/v2/models"
)

func List() ([]models.Task, error) {
	var task []models.Task

	readfile, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(readfile, &task); err != nil {
		return nil, err
	}

	return task, nil
}

func ListByfilter(status string) ([]models.Task, error) {
	var task []*models.Task

	readfile, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(readfile, &task); err != nil {
		return nil, err
	}

	var data []models.Task
	if status != "" {
		for _, v := range task {
			if v.Status == status {
				data = append(data, models.Task{
					Id:          v.Id,
					Description: v.Description,
					Status:      v.Status,
					CreatedAt:   v.CreatedAt,
					UpdatedAt:   v.UpdatedAt,
				})
			}
		}
	}

	return data, nil
}

// add task
func AddTask(desc string) (string, error) {
	var task []*models.Task
	id := 1

	openfile, err := os.OpenFile("data.json", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return "", err
	}

	file, err := os.ReadFile("data.json")
	if err != nil {
		return "", err
	}

	erer := json.Unmarshal(file, &task)
	if erer != nil {
		id = 1
	}

	if len(task) > 0 {
		lastdata := task[len(task)-1]
		id = lastdata.Id + 1
	}

	tasks := &models.Task{
		Id:          id,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now(),
	}

	task = append(task, tasks)

	databyte, err := json.Marshal(&task)
	if err != nil {
		return "", err

	}

	defer openfile.Close()

	if _, err := openfile.Write(databyte); err != nil {
		return "", err

	}
	msg := fmt.Sprintf("Task added successfully (ID: %d)", id)
	return msg, nil
}

// update task
func UpdateTask(id, desc string) (string, error) {
	var task []*models.Task

	ids, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	readfile, err := os.ReadFile("data.json")
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(readfile, &task); err != nil {
		return "", err
	}

	found := false
	for _, v := range task {
		if v.Id == ids {
			found = true
			v.Description = desc
			v.UpdatedAt = time.Now()
		}
	}

	if !found {
		return "", fmt.Errorf("task with ID %d not found", ids)
	}

	jsondata, err := json.Marshal(task)
	if err != nil {
		return "", err
	}

	openfile, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return "", err
	}
	defer openfile.Close()

	if _, err := openfile.Write(jsondata); err != nil {
		return "", err
	}

	msg := fmt.Sprintln("Task Updated successfully")
	return msg, nil
}

func DeleteTask(id string) (string, error) {
	var task []*models.Task

	ids, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	readfile, err := os.ReadFile("data.json")
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(readfile, &task); err != nil {
		return "", err
	}

	var data []*models.Task
	found := false
	for _, v := range task {
		if v.Id != ids {
			found = true
			data = append(data, &models.Task{
				Id:          v.Id,
				Description: v.Description,
				Status:      v.Status,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
			})
		}
	}

	if found {
		return "", fmt.Errorf("task with ID %d not found", ids)
	}

	datajson, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	openfile, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return "", err
	}
	defer openfile.Close()

	if _, err := openfile.Write(datajson); err != nil {
		return "", err
	}

	msg := fmt.Sprintln("Task delete successfully")
	return msg, nil
}

// mark in progress
func MarkInProgress(id string) (string, error) {
	var task []*models.Task

	ids, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	readfile, err := os.ReadFile("data.json")
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(readfile, &task); err != nil {
		return "", err
	}

	found := false
	for _, v := range task {
		if v.Id == ids {
			found = true
			v.Status = "in-progress"
			v.UpdatedAt = time.Now()
		}
	}

	if !found {
		return "", fmt.Errorf("task with ID %d not found", ids)
	}

	jsondata, err := json.Marshal(task)
	if err != nil {
		return "", err
	}

	openfile, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return "", err
	}
	defer openfile.Close()

	if _, err := openfile.Write(jsondata); err != nil {
		return "", err
	}

	msg := fmt.Sprintln("Task Updated successfully")
	return msg, nil
}

func MarkDone(id string) (string, error) {
	var task []*models.Task

	ids, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	readfile, err := os.ReadFile("data.json")
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(readfile, &task); err != nil {
		return "", err
	}

	found := false
	for _, v := range task {
		if v.Id == ids {
			found = true
			v.Status = "done"
			v.UpdatedAt = time.Now()
		}
	}

	if !found {
		return "", fmt.Errorf("task with ID %d not found", ids)
	}

	jsondata, err := json.Marshal(task)
	if err != nil {
		return "", err
	}

	openfile, err := os.OpenFile("data.json", os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return "", err
	}
	defer openfile.Close()

	if _, err := openfile.Write(jsondata); err != nil {
		return "", err
	}

	msg := fmt.Sprintln("Task Updated successfully")
	return msg, nil
}
