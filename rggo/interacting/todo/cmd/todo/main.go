package main

import (
	"flag"
	"fmt"
	"os"

	"pragprog.com/rggo/interacting/todo"
)

//hardcoding the file name

const todoFileName = ".todo.json"

func main(){
	//Parsing command line flags 
	task := flag.String("task", " ", "Task to be included in the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()
	//using the address operator & to extract the address of an empty instance of todo.List
	l := &todo.List{}

	//use the get method to read todo items from file

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}

	//using flags 
	switch {
	case *list:
		//list current todo items 
		for _,item := range *l {
			//exclude the completed items from the output 
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		//complete the given item 
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		//save the new list 
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	

	case *task != " ":
		//Add the task 
		l.Add(*task)

		//save the new list 
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		//invalid flag provided 
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)
	}




}


// 	//decide what to do based on the number of arguments provided 
// 	switch {
// 	//For no extra arguments print the list 
// 	case len(os.Args) == 1:
// 		// List current todo items 
// 		for _, item := range *l {
// 			fmt.Println(item.Task)
// 		}
// 	//concatenate all provided arguments with a space and add to the list as an item
// 	default:
// 		//concatenate all arguments with a space 
// 		//the items we want to add to the list are passed in as arguments. we start from [1:] because [0] is the path to the command in the args slice
// 		//so we concatenate the items that come after the command 
// 		item := strings.Join(os.Args[1:], " ")
// 		//add the item to the list 
// 		l.Add(item)

// 		//save the new list 
// 		if err := l.Save(todoFileName); err != nil{
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(1)
// 		}

// 	}
// }

//create a json todo file
//Get a list 
//read existing items from the file and save to list
//if only the command was put in, then print out the items in the list 
//else if items were added after the command, add them to the list 
//save the new list to the file