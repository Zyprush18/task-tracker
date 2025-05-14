package main

import (
	"fmt"
	"os"

	"github.com/Zyprush18/task-tracker/service"

)

func main() {
	argument:= os.Args

	if argument[1] != "" {
		switch argument[1] {
		case "add":
			msg,err := service.AddTask(argument[2])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(msg)
		case "update":
			msg,err := service.UpdateTask(argument[2], argument[3])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(msg)
		default:
			fmt.Println("Unknown")
		}
	} else{
		fmt.Println("please enter according to the order")
	}
	


}