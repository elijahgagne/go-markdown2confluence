package main

import (
	"github.com/justmiles/go-markdown2confluence/cmd"
)

// version of markdown2confluence. Overwritten during build
var version = "3.1.1-2020_05_27"

func main() {
	cmd.Execute(version)
}
