package main

import (
	"github.com/foxboolean/go_tour/go_tools/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd Execute err: %v", err)
	}
}
