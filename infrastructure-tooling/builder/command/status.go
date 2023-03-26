/*
 * Copyright 2020 Solus Project <copyright@getsol.us>
 *
 * Author: Bryan T. Meyers
 */

package main

import (
	"github.com/jmoiron/sqlx"
	"time"
)

// Change the status of a job
func status(db *sqlx.DB, req Request) Response {
	// pivot on status
	switch req.Job.Status {
	case "FAILED", "OK":
		// Mark completed
		req.Job.Finished.String = time.Now().Format(time.ANSIC)
		req.Job.Finished.Valid = true
		fallthrough
	default:
		// Update record in the DB
		if _, err := db.NamedExec(UpdateJob, &req.Job); err != nil {
			// DB/IO error
			return Response{
				Message: "Failed to update status",
				Error:   err.Error(),
			}
		}
	}
	// Success
	return Response{
		Message: "Successfully updated job status",
		Job:     req.Job,
	}
}
