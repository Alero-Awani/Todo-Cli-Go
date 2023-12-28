package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"pragprog.com/rggo/interacting/todo"
)

//hardcoding the file name

var todoFileName = ".todo.json"



func main(){

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for the Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2023\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information")
		flag.PrintDefaults()
	}
	//Parsing command line flags 
	add := flag.Bool("add",false, "Add task to the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("delete", 0, "Delete an item from the list")

	flag.Parse()

	//Check if the user defined the ENV VAR for the custom file name
	if os.Getenv("TODO_FILENAME") != ""{
		todoFileName = os.Getenv("TODO_FILENAME")
	}

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
		//List current todo item by implementing fmt.Stringer interface
		fmt.Print(l)
		// //list current todo items 
		// for _,item := range *l {
		// 	//exclude the completed items from the output 
		// 	if !item.Done {
		// 		fmt.Println(item.Task)
		// 	}
		// }
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
	

	case *add:
		//Add the task 
		//when any argument(excluding flags) are provided, they will be used as the new task

		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		//save the new list 
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		// delete the task
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		//save the new list 
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	default:
		//invalid flag provided 
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)
	}
}


//getTask function decides where to get the description for a new 
//task from , args or STDIN

func getTask(r io.Reader, args ...string)(string, error){
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(),nil

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