package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

//This tool previews Markdown files locally, using a web browser

const (
	header = `<!DOCTYPE html>
<html>
 	<head>
	 <meta http-equiv="content-type" content="text/html; charset=utf-8">
	 <title>Markdown Preview Tool</title>
 	</head>
 	<body>
`

	footer = `
 	</body>
</html>
`
)


func main() {
	//Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")

	flag.Parse()

	//If the user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//this function coordinates the execution of the remaining functions. 
//the function reads the markdown file into a [] bytes using ReadFile
func run(filename string, out io.Writer, skipPreview bool) error {
	//Read all the data from the input file and check for errors
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	//responsible for converting markdown to html
	htmlData := parseContent(input)

	//Create temporary fila and check for errors 
	temp, err := ioutil.TempFile("./tmp", "mdp*.html")
	if err != nil {
		return err
	}

	if err := temp.Close(); err != nil {
		return err
	}

	outName := temp.Name()
	fmt.Fprintln(out, outName)


	// // this is the output html filename 
	// outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	// fmt.Println(outName)

	// this saves the html content to a file//this returns a POTENTIAL ERROR which the function run() also returns as its error
	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}
	defer os.Remove(outName)

	return preview(outName)
}



//thus fucntion receives a slice of bytes representing the content of the Markdown file and returns another []bytes
//with the converted content as HTML
func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	//the block of code above generated HTMl that constitutes the body
	//Now combine this body with the header and footer to generate the complete HTML content

	//create a buffer of bytes to write to file
	var buffer bytes.Buffer

	//Write html to bytes bufefer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes() 

}


func saveHTML(outFname string, data []byte) error {
	//Write the bytes to the file
	return ioutil.WriteFile(outFname, data, 0644)
}


//note that in the exec command, The first parameter is the program to be run; 
//the other arguments are parameters to the program.
func preview(fname string) error {
	cName := ""
	cParams := []string{}

	//Define executable based on OS
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}
	// Append file name to parameters slice 
	cParams = append(cParams, fname)

	//Locate executable in PATH
	cPath, err := exec.LookPath(cName)

	if err != nil {
		return err 
	}

	// Open the file using the default program
	err = exec.Command(cPath, cParams...).Run()

	//Give the browser some time to open the file before deleting it 
	time.Sleep(2 * time.Second)
	return err 

}



//the black friday package converts markdown to html but doesnt sanitize its output to
//prevent malicious content. 
//tha bluemonday package sanitizes this output


//BUFFER - The buffer name itself clarifies its purposes; it allows us to give buffer storage where we can store 
//some data and access the same data when needed. 