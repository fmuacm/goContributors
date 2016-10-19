package main

/*
 * This is a simple golang console application that asks for your name as input
 * and appends your name to the end of the README.md file in this repository.
 * Author: trey.watford@acstechnologies.com
 * Date: 10/17/16
 */

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

//use newline character as line delimiter
const inputDelimiter = '\n'

func main() {
	//prompt for user input
	fmt.Print("Please enter your name: ")

	//create input reader to read line from stdin
	reader := bufio.NewReader(os.Stdin)

	//read from stdin ending in newline
	in, err := reader.ReadString(inputDelimiter)
	if err != nil {
		fmt.Println(err)
		return
	}

	//remove CRLF in windows strings and LF in unix strings
	if runtime.GOOS == "windows" {
		in = strings.Replace(in, "\r\n", "", -1)
	} else {
		in = strings.Replace(in, "\n", "", -1)
	}

	fmt.Printf("You entered %s, appending to file.\n", in)

	//open file README.md in append or write only mode
	appendFile, err := os.OpenFile("README.md", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	//defer file close to end of main function
	defer appendFile.Close()

	//append to file with github markdown for bullets *
	if _, err = appendFile.WriteString("* " + in + "\n"); err != nil {
		panic(err)
	}

	//inform user and exit
	fmt.Printf("%s, appended to file README.md.\n", in)
}
