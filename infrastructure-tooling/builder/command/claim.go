/*
 * Copyright 2020 Solus Project <copyright@getsol.us>
 *
 * Author: Bryan T. Meyers
 */

package main

import (
	"github.com/jmoiron/sqlx"
)

// Get the next available job
func claim(db *sqlx.DB) Response {
	j := Job{}
	// Find a job
	if err := db.Get(j, GetJob); err != nil {
		// No jobs, or some other error
		return Response{
			Message: "No jobs available",
			Error:   err.Error(),
		}
	}
	// Mark the job as claimed
	j.Status = "CLAIMED"
	if _, err := db.NamedExec(UpdateJob, &j); err != nil {
		return Response{
			Message: "Failed to claim job",
			Error:   err.Error(),
		}
	}
	// Return the job information
	return Response{
		Message: "Successfully claimed job",
		Job:     j,
	}
}
