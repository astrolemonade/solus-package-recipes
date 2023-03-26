/*
 * Copyright 2020 Solus Project <copyright@getsol.us>
 *
 * Author: Bryan T. Meyers
 */

package main

import (
	"database/sql"
)

// Job represents a scheduled package build
type Job struct {
	ID       int64          `db:"ID"`
	Pkg      string         `db:"PKG"`
	Tag      string         `db:"TAG"`
	Status   string         `db:"STATUS"`
	Builder  string         `db:"BUILDER"`
	Finished sql.NullString `db:"FINISHED"`
}

// JobSchema is the table structure of for the jobs database
const JobSchema = `
CREATE TABLE IF NOT EXISTS JOB_QUEUE(
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    PKG TEXT,
    TAG TEXT,
    STATUS TEXT,
    BUILDER TEXT,
    FINISHED TEXT
)
`

// InsertJob adds a new Job to the DB
const InsertJob = `
INSERT INTO JOB_QUEUE(
    PKG, TAG, STATUS, BUILDER
) VALUES (
    :PKG, :TAG, "UNCLAIMED", :BUILDER
)
`

// GetJob requests the first unclaimed job
const GetJob = `SELECT * FROM JOB_QUEUE WHERE STATUS="UNCLAIMED" LIMIT 1`

// UpdateJob is used to set the status of a job
const UpdateJob = `UPDATE JOB_QUEUE SET (STATUS=:STATUS,FINISHED=:FINISHED) WHERE ID=:ID`

// GetLatestJobs returns the 30 most recent jobs
const GetLatestJobs = "SELECT * FROM JOB_QUEUE ORDER BY ID DESC LIMIT 30"
