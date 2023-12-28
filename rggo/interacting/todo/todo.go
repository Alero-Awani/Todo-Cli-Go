package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type List []item//a list of struct item

// add method creates a new todo item and appends it to the list
func (l *List) Add(task string){
	t := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)

}

//Complete method marks a todo item as completed by setting Done to true and completed at to the current time

func (l *List) Complete(i int) error {
	ls := *l //storing the derefrerrenced list so it can be used because what we input as the receiver is a pointer to the list
	//checking if the item number falls within the range of the list 
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	//if it does then we will carry out the code below //adjusted for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil

}

//DELETE method 
func (l *List) Delete(i int) error {
	ls := *l 

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Items %d does not exist", i)
	}
	// i is the id of the item on the list 
	*l = append(ls[:i-1], ls[i:]...)
	return nil //because our only return value in the function definition is an error, so if everything works out fine then no error will be returned 
}

//Save method which encodes the list as json and saves it using the provided file name 
func (l *List) Save(filename string) error{
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

//Get method opens the provided file name, decodes the json data and parses it into a List

func(l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		//checks if the file does not exist
		if errors.Is(err, os.ErrNotExist){
			return nil
		}
		return err
	}

	//
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

