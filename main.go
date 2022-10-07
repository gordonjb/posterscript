package main

import "github.com/gordonjb/posterscript/cmd"

var (
	// these will be replaced by goreleaser
	version = "v0.0.0"
	date    = "0001-01-01T00:00:00Z"
	commit  = "0000000"
)

func main() {
	cmd.Execute(version, date, commit)
}
