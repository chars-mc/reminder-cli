package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chars-mc/reminder-cli/client"
)

var (
	backendURI = flag.String("backend", "http://localhost:3000", "Backend API URL")
	helpFlag   = flag.Bool("help", false, "Display the help message")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendURI)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error: %s", err)
		os.Exit(2)
	}
}
