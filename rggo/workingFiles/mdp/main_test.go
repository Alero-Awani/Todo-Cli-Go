package main

import ( "bytes"
  "io/ioutil"
  "os"
  "testing"
  "strings"
)


const (
	inputFile = "./testdata/test1.md"
	// resultFile = "test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)


//recall the parsefunction takes in the markdown file and converts it to html
func TestParseContent(t *testing.T) {

	//read in the mark down file
	input, err := ioutil.ReadFile(inputFile)

	if err != nil {
		t.Fatal(err)
	}

	result := parseContent(input)

	expected, err := ioutil.ReadFile(goldenFile)

	if err != nil {
		t.Fatal(err)
	}

	//bytes.Equal() compares two slices of bytes
	if !bytes.Equal(expected, result) {
		t.Logf("golden len: %d", len(expected))
		t.Logf("result len: %d", len(result))

		t.Logf("golden: \n%s\n", expected)
		t.Logf("result: \n%s\n", result)
		t.Error("Result content does not match golden file")
	}
}

//An integrated test case that tests the run fucntion
func TestRun(t *testing.T) {
	var mockStdOut bytes.Buffer

	if err := run(inputFile, &mockStdOut); err != nil {
		t.Fatal(err)
	}

	//here we are using bytes.Buffer to capture the output filename
	//note that bytes.Buffer satisfies the io.Write interface using a pointer receiver which is why we passed the address of the buffer to run
	resultFile := strings.TrimSpace(mockStdOut.String())


	result, err := ioutil.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, result) {
		t.Logf("golden: \n%s\n", expected)
		t.Logf("result: \n%s\n", result)
		t.Error("Result content does not match golden file")
	}

	os.Remove(resultFile)
}




//GOLDEN FILES - with goldenfiles, the expected results are saved into files that are 
//loaded during the tests for validating the actual output