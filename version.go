package version

import (
	"fmt"
)

var (
	Name           string
	Version        string
	CommitID       string
	LastCommitTime string
)

// Show print version
func Show() {
	fmt.Println(String())
}

// String return version
func String() string {
	return fmt.Sprintf("Program: %s\nVersion: %s\nCommitID: %s\nLastCommitTime: %s\n",
		Name, Version, CommitID, LastCommitTime)
}
