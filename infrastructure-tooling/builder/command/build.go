/*
 * Copyright 2020 Solus Project <copyright@getsol.us>
 *
 * Author: Bryan T. Meyers
 */

package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// queue up a new build job
func build(db *sqlx.DB, req Request) Response {
	j := req.Job
	// Create the new entry
	result, err := db.NamedExec(InsertJob, j)
	if err != nil {
		// Probably a DB error
		return Response{
			Message: "Failed to insert Job",
			Error:   err.Error(),
			Job:     j,
		}
	}
	// Get the job ID
	id, err := result.LastInsertId()
	if err != nil {
		// Probably a DB error too
		return Response{
			Message: "Failed to get Job ID",
			Error:   err.Error(),
			Job:     j,
		}
	}
	// Success
	j.ID = id
	return Response{
		Message: fmt.Sprintf("Building: %s (%s)\n", j.Pkg, j.Tag),
		Job:     j,
	}
}
