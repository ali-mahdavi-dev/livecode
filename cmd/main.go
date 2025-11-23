package main

import (
	"log"

	"livecode/cmd/command"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	command.Execute()
}
