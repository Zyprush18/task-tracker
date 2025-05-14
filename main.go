package main

import (
	"fmt"
	"os"

	"github.com/Zyprush18/task-tracker/service"
)

func main() {
	argument := os.Args

	if len(argument) > 1 {
		switch argument[1] {

		case "list":
			if len(argument) > 2 {
				switch argument[2] {
				case "todo", "in-progress", "done":
					data, err := service.ListByfilter(argument[2])
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					fmt.Printf("%+v\n", data)
				default:
					fmt.Println("Unknown status filter. Use: todo, in-progress, or done.")
				}
			} else {
				data, err := service.List()
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Printf("%+v\n", data)
			}

		case "add":
			if len(argument) > 2 {
				msg, err := service.AddTask(argument[2])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(msg)
			} else {
				fmt.Println("Missing description for add")
			}

		case "update":
			if len(argument) > 3 {
				msg, err := service.UpdateTask(argument[2], argument[3])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(msg)
			} else {
				fmt.Println("Missing ID or description for update")
			}

		case "delete":
			if len(argument) > 2 {
				msg, err := service.DeleteTask(argument[2])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(msg)
			} else {
				fmt.Println("Missing ID for delete")
			}

		case "mark-in-progress":
			if len(argument) > 2 {
				msg, err := service.MarkInProgress(argument[2])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(msg)
			} else {
				fmt.Println("Missing ID to mark in-progress")
			}

		case "mark-done":
			if len(argument) > 2 {
				msg, err := service.MarkDone(argument[2])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(msg)
			} else {
				fmt.Println("Missing ID to mark done")
			}

		default:
			fmt.Println("Unknown command")
		}
	} else {
		fmt.Println("Please enter a command")
	}

}
