/*
 * Copyright 2020 Solus Project <copyright@getsol.us>
 *
 * Author: Bryan T. Meyers
 */

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

const port = 798
const user = "build"
const host = "getsol.us"

// Handle sending the request
func send(stdin io.WriteCloser, req Request) {
	defer stdin.Close()
	enc := json.NewEncoder(stdin)
	if err := enc.Encode(&req); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// This is not a daemon, Damon, or demon
func main() {
	var req Request
	var resp Response
	var err error
	// Check for subcommand
	if len(os.Args) < 2 {
		println("Command is required")
		os.Exit(1)
	}
	// Setup the request
	req.Method = os.Args[1]
	switch req.Method {
	case "build":
		// Check number of args
		if len(os.Args) != 4 {
			fmt.Println("Not enough arguments")
			os.Exit(1)
		}
		// Setup job
		req.Job = Job{
			Pkg: os.Args[2],
			Tag: os.Args[3],
		}
	default:
		fmt.Printf("'%s' is not a valid method.\n", req.Method)
		os.Exit(1)
	}
	// Setup the command
	cmd := exec.Command("ssh", "-4", "-p", strconv.Itoa(port), fmt.Sprintf("%s@%s", user, host))
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	go send(stdin, req)
	// Run the command
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Parse the response
	if err = json.Unmarshal(out, &resp); err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Response: %s\n", string(out))
		os.Exit(1)
	}
	// Check for failure
	if len(resp.Error) > 0 {
		fmt.Printf("%s: %s\n", resp.Message, resp.Error)
		os.Exit(1)
	}
	// Print the response
	fmt.Println(resp.Message)
	resp.Job.Print()
}
