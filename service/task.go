package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Zyprush18/task-tracker/models"
)

// func init()  {
// 	_,err := os.Stat("data.json")
// }

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
	}

	if len(task) != 0 {
		id = len(task) + 1
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
