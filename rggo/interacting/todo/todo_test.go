package todo_test

import (
	"testing"
	"io/ioutil"
	"os"

	"pragprog.com/rggo/interacting/todo"
)

//creating a test to ensure we can add items to the list 
func TestAdd(t *testing.T) {
	l := todo.List{} //importing the list type from the todo package and creating the l variable of typoe todo.List{}

	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

}

//This tests the complete method of the list type 

func TestComplete(t *testing.T) {
	l := todo.List{}

	//a task needs to be added first 
	taskName := "New Task"

	l.Add(taskName)

	//check if task was properly added to the list 
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	//if l[0].Done results to true, then it will throw an error that new task is not meant to be Done(true) because the default is actually false
	if l[0].Done {
		t.Errorf("New task should not be completed")
	}

	//call the complete function to complete the task
	l.Complete(1)

	//after calling the complete method, l becomes true, so if its not true(false) then it will throw an error that its meant to be true because the co
	//complete function has been called
	if !l[0].Done {
		t.Errorf("New task should be completed")
	}

}

//TestDelete tests the Delete method of the list type 

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New task 1",//0
		"New task 2",//1
		"New task 3",//2
	}

	for _, v := range tasks{
		l.Add(v)
	}

	
	if l[0].Task != tasks[0]{
		t.Errorf("Expected %q, got %q instead",tasks[0], l[0].Task)
	}

	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead", 2, len(l))
	}

	//because 2 has been deleted, 
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)

	}
}

//TestSaveget which tests the save and get method of the list type

 
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"

	l1.Add(taskName)

	if l1[0].Task != taskName{
		t.Errorf("Expected %q, got %q instead",taskName, l1[0].Task)
	}

	//create temp file with ioutil

	tf, err := ioutil.TempFile("","")

	if err != nil{
		t.Fatalf("Error creating temp file: %s", err)
	}
	//remove the temp file after we are done
	defer os.Remove(tf.Name())

	//saving the contents of the list to the file
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	//getting the json content of the file and saving it to a list 
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	//chek if the lists of both instances are the same 

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match Task %q", l1[0].Task, l2[0].Task)
	}
}