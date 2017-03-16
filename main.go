package main

import (
	"log"

	"fmt"

	"github.com/sascha-andres/gitbranch/app"
)

func main() {
	branches, err := app.GetBranches("git@gitssh.syzygy.de:deuba/ctswebtools.git")
	if err != nil {
		log.Fatal(err)
	}
	for _, branch := range branches {
		fmt.Println(branch.Value)
	}
}
