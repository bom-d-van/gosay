// Package gosay is simple wrapper for Mac OS X say command. It can tell whether a text is Japanese or English.
//
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	pidFile := filepath.Join(usr.HomeDir, ".gosay-say-pid")
	log.SetFlags(log.Lshortfile)
	pidStr, err := ioutil.ReadFile(pidFile)
	if err != nil {
		log.Println(err)
	} else if pid, err := strconv.Atoi(string(pidStr)); err == nil {
		ps, err := os.FindProcess(int(pid))
		if err != nil {
			log.Println(err)
		} else if err := ps.Signal(syscall.Signal(0)); err == nil {
			if err := ps.Kill(); err != nil {
				log.Println(err)
			} else {
				return
			}
		} else {
			// log.Println(err)
		}
	} else {
		log.Println(err)
	}

	if len(os.Args) == 1 {
		println("expect input")
		os.Exit(1)
	}
	words := strings.Join(os.Args[1:], " ")
	c := []rune(words)[0]
	var args []string
	if 0 <= c && c <= 127 {
		args = []string{words}
	} else {
		args = []string{"-v", "Otoya", words}
	}
	cmd := exec.Command("say", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Printf("error: %s\n", err)
		os.Exit(1)
	}
	file, err := os.Create(pidFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	file.WriteString(fmt.Sprint(cmd.Process.Pid))
	file.Close()
}
