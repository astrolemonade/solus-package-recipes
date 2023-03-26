/*
 * Copyright 2020 Solus Project <copyright@getsol.us>
 *
 * Author: Bryan T. Meyers
 */

package main

import (
	"github.com/jmoiron/sqlx"
	"html/template"
	"os"
)

// Index is a template for the build page
const Index = `
<html>
<head>
<title>Package Build Status</title>
<style>

/*
Taken from: http://johnsardine.com/freebies/dl-html-css/simple-little-tab/
Table Style - This is what you want
------------------------------------------------------------------ */
table a:link {
        color: #26292B;
        text-decoration:none;
}
table a:visited {
        color: #999999;
        font-weight:bold;
        text-decoration:none;
}
table a:active,
table a:hover {
        color: #1A559C;
        text-decoration:underline;
}
table {
        font-family:Roboto;
        color: #26292B;
        font-size:12px;
        background:#F5F5F5;
        margin:20px;
        border: 1px solid #26292B;
        border-radius: 2px;
}
table th {
        padding:21px 25px 22px 25px;

        color:#F5F5F5;
        background: #26292B;
}
table th:first-child {
        text-align: left;
        padding-left:20px;
}
table tr:first-child th:first-child {
        -moz-border-radius-topleft:2px;
        -webkit-border-top-left-radius:2px;
        border-top-left-radius:2px;
}
table tr:first-child th:last-child {
        -moz-border-radius-topright:2px;
        -webkit-border-top-right-radius:2px;
        border-top-right-radius:2px;
}
table tr {
        padding-left:20px;
        text-align: left;
}
table td:first-child {
        text-align: left;
        padding-left:20px;
        border-left: 0;
}
table td {
    padding:18px;
        border-bottom:1px solid #e0e0e0;
        background: #fafafa;
        text-align: left;
}
table tr:nth-child(2n) td {
        background: #eee;
}
table tr:last-child td {
        border-bottom:0;
}
table tr:last-child td:first-child {
        -moz-border-radius-bottomleft:2px;
        -webkit-border-bottom-left-radius:2px;
        border-bottom-left-radius:2px;
}
table tr:last-child td:last-child {
        -moz-border-radius-bottomright:2px;
        -webkit-border-bottom-right-radius:2px;
        border-bottom-right-radius:2px;
}
table tr:hover td {
        background: #f2f2f2;
}

.FAILED {
    color: red;
}
.BUILDING {
    color: orange;
}
.UNCLAIMED {
    color: grey;
}
.CLAIMED {
    color: blue;
}
.OK {
    color: green;
}
</style>
</head>
<body>
    <center><img src="/logo.png" alt="Solus Build System" /></center>
    <br />
    <center>
    <table class="tablesorter" cellspacing='0' width="80%">
        <thead>
            <tr>
                <th># ID</th>
                <th>Package</th>
                <th>Built by</th>
                <th>Finished</th>
                <th>Status</th>
            </tr>
        </thead>
        <tbody>
        {{range .}}
            <tr>
                <td>{{.ID}}</td>
                {{if or (eq .Status "FAILED") (eq .Status "OK")}}
                <td>
                    <a href="/logs/{{.Tag}}.log">
                        {{.Pkg}} ({{.Tag}})
                    </a>
                </td>
                {{else}}
                <td>{{.Pkg}}</td>
                {{end}}
                <td>
                    <a href="https://dev.getsol.us/source/{{.Pkg}}/browse/master/;{{.Tag}}">
                        {{.Builder}}
                    </a>
                </td>
                <td>{{.Finished.String}}</td>
                <td class="{{.Status}}">{{.Status}}</td>
            </tr>
        {{end}}
        </tbody>
    </table>
    <br />
    <div><small>Copyright &copy; 2015-2020 Solus Project. The Solus logo is Copyright &copy; 2016-2020 Solus Project.</small></div>
    </center>
</body>
</html>
`

// Generate the HTML build page
func genHTML(db *sqlx.DB, home string) Response {
	// Get the most recent jobs
	jobs := []Job{}
	if err := db.Select(&jobs, GetLatestJobs); err != nil {
		// probably a DB error
		return Response{
			Message: "Failed to get latest jobs",
			Error:   err.Error(),
		}
	}
	// Parse the template
	tmpl, err := template.New("index").Parse(Index)
	if err != nil {
		// probably a typo or something stupid
		return Response{
			Message: "Failed to parse index template",
			Error:   err.Error(),
		}
	}
	// Open a new HTML file
	index, err := os.Create(home + "/index.html")
	if err != nil {
		// IO/permissions error
		return Response{
			Message: "Failed to create index file",
			Error:   err.Error(),
		}
	}
	defer index.Close()
	// Generate the build page
	if err := tmpl.Execute(index, jobs); err != nil {
		// IO/template error
		return Response{
			Message: "Failed to populate index file",
			Error:   err.Error(),
		}
	}
	// Success
	return Response{
		Message: "Index Updated successfully",
	}
}
