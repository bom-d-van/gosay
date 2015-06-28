// Package gosay is simple wrapper for Mac OS X say command. It can tell whether a text is Japanese or English.
//
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		println("expect input")
		os.Exit(1)
	}
	words := os.Args[1]
	var output []byte
	var err error
	c := []rune(words)[0]
	if 0 <= c && c <= 127 {
		output, err = exec.Command("say", words).CombinedOutput()
	} else {
		output, err = exec.Command("say", "-v", "Otoya", words).CombinedOutput()
	}
	if err != nil {
		fmt.Printf("error (%s): %s\n", err, output)
	}
}
