package service

import (
	"encoding/json"
	"fmt"
	"os"	
	"strconv"
	"time"

	"github.com/Zyprush18/task-tracker/models"
)

// add task
func AddTask(desc string) (string, error) {
	var task []*models.Task
	var id int

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
		fmt.Println(err)
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
	return msg,nil
}

// update task
func UpdateTask(id,desc string) error {
	var task []*models.Task

	ids, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	readfile, err := os.ReadFile("data.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(readfile, &task); err != nil{
		return err
	}

	fmt.Println(task)
	for _, v := range task {
		if v.Id == ids {
			v.Description = desc
			v.UpdatedAt = time.Now()
		}
	}


	jsondata, err := json.Marshal(task)
	if err != nil {
		return err
	}

	openfile, err := os.OpenFile("data.json",os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer openfile.Close()

	if _,err:= openfile.Write(jsondata);err != nil {
		return err
	}
	return nil
}
