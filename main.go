package main

import (
	"fmt"

	"carl/cmd"
)

var (
	buildVersion string
	buildTime    string
	buildCommit  string
)

func main() {
	err := cmd.Run(buildVersion, buildTime, buildCommit)
	if err != nil {
		fmt.Println(err.Error())
	}
}
