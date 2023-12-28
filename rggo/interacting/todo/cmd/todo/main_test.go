package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	

)

var (
	binName = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool.....")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up..")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}


func TestTodoCLI(t *testing.T) {
	//define the task name 
	task := "test task number 1"
	// num := 1

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	//this creates the path to the command todo
	cmdPath := filepath.Join(dir, binName)

	//Create a new test that ensures the tool can add a new task by using t.Run
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-task", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T){
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := fmt.Sprintf("  1: %s\n", task)

		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})

	//test to ensure that the tool deletes a task
	// t.Run("DeleteTasks", func(t *testing.T){
	// 	cmd := exec.Command(cmdPath, "-delete", num)
	// 	if err := cmd.Run(); err != nil {
	// 		t.Fatal(err)
	// 	}
	// })
}

//what happens here is the code in test main runs first, then when m.Run is called , the tests then run 


//these tests above are integration tests which are used to test the user interface, they can be used to test the expected output 
//in line with the api(todo.go)