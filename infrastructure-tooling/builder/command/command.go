/*
 * Copyright 2020 Solus Project <copyright@getsol.us>
 *
 * Author: Bryan T. Meyers
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// This is not a daemon, Damon, or demon
func main() {
	var resp Response
	var db *sqlx.DB
	var err error
	// Check for builder
	if len(os.Args) < 2 {
		println("Builder name is required")
		os.Exit(1)
	}
	builder := os.Args[1]
	// Read request
	var req Request
	var home string
	dec := json.NewDecoder(os.Stdin)
	if err = dec.Decode(&req); err != nil {
		// Malformed request
		resp = Response{
			Message: "Failed to parse request",
			Error:   err.Error(),
		}
		goto RESPOND
	}
	// Set builder
	req.Job.Builder = builder
	// Get the builder's home directory
	if home, err = os.UserHomeDir(); err != nil {
		// User / IO error
		resp = Response{
			Message: "Failed to get home directory",
			Error:   err.Error(),
		}
		goto RESPOND
	}
	// Connect to DB
	if db, err = sqlx.Open("sqlite3", "file:"+home+"/build_database?cache=shared"); err != nil {
		resp = Response{
			Message: "Failed to open DB",
			Error:   err.Error(),
		}
		goto RESPOND
	}
	defer db.Close()
	// Limit connections for safety
	db.SetMaxOpenConns(1)
	// Make sure the table exists
	db.MustExec(JobSchema)
	// Find a command
	switch req.Method {
	case "build":
		resp = build(db, req)
	case "claim":
		resp = claim(db)
	case "status":
		resp = status(db, req)
	default:
		// Not a command
		resp = Response{
			Message: "Could not complete request",
			Error:   fmt.Sprintf("'%s' is an invalid method", req.Method),
		}
	}
RESPOND:
    // Update the index page
	if resp2 := genHTML(db, home); len(resp.Error) > 0 {
        resp = resp2
    }
	// Send the response to the client
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&resp); err != nil || len(resp.Error) > 0 {
		os.Exit(1)
	}
}
