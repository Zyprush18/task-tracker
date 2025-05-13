package main

import (
	// "fmt"
	"fmt"
	"os"

	"github.com/Zyprush18/task-tracker/service"
)

func main() {
	argument:= os.Args

	switch argument[1] {
	case "add":
		msg,err := service.AddTask(argument[2])
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(msg)

	default:
		fmt.Println("Unknown")
	}


}