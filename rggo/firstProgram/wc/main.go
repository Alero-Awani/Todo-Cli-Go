package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main(){
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b",false,"Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}


func count(r io.Reader, countLines bool, countBytes bool) (int int){

	//the scanner reads data delimited by new spaces or lines by default
	scanner := bufio.NewScanner(r)

	//here we instruct the scanner to read words instead by using the split function
	//if countLines flag is not set, we want to count words
	if !countLines && !countBytes{
		scanner.Split(bufio.ScanWords)
	}
	//if the its true then run this
	if countBytes{
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0

	for scanner.Scan(){	
		wc++
	}

	return wc

}


//test program using the executable by passing an input string with 
// echo "My first command line tool with Go" | ./wc

//count number of lines in the main.go file with 
//cat main.go | wc -l 