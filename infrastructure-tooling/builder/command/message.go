/*
 * Copyright 2020 Solus Project <copyright@getsol.us>
 *
 * Author: Bryan T. Meyers
 */

package main

// Request is used to ask make requests
type Request struct {
	Method string
	Job    Job `json:",omitempty"`
}

// Response is given as the result of a Request being executed
type Response struct {
	Message string
	Error   string `json:",omitempty"`
	Job     Job    `json:",omitempty"`
}
