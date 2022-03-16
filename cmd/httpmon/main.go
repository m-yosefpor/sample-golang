package main

import (
	"log"

	_ "go.uber.org/automaxprocs"

	"github.com/m-yosefpor/httpmon/internal/cmd"
)

func main() {
	root := cmd.NewRootCommand()
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
