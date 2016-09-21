package main

import (
	"fmt"
	"github.com/pazams/mdgofmt"
	"io/ioutil"
	"os"
)

var usage = `
Usage example: 

cat <input-file> | mdgofmt-cli > <output-file>

`

func printUsage() {
	fmt.Fprint(os.Stderr, usage)
}

func main() {
	// Read
	stat, err := os.Stdin.Stat()
	checkErr(err)
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		usageAndExit("No data in STDIN")
	}
	bytes, err := ioutil.ReadAll(os.Stdin)
	checkErr(err)

	// Formats golang code blocks inside markdown
	out, err := mdgofmt.Format(bytes)
	checkErr(err)
	if len(out) == 0 {
		usageAndExit("Could not format the input")
	}

	// Write
	_, err = os.Stdout.WriteString(string(out))
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		usageAndExit(err.Error())
	}
}

func usageAndExit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	printUsage()
	os.Exit(1)
}
