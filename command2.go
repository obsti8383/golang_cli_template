package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func Command2(arguments []string, configuration Configuration, errorLogger *log.Logger, debugLogger *log.Logger) int {
	// eval flags
	var help bool
	var verbose bool
	var pages int
	flags := flag.NewFlagSet("command2", flag.ContinueOnError)
	flags.BoolVar(&help, "help", help, "Show this help message")
	flags.BoolVar(&help, "h", help, "")
	flags.BoolVar(&verbose, "v", verbose, "Show verbose logging.")
	flags.IntVar(&pages, "count", pages, "Maximum number of entries to fetch")
	err := flags.Parse(arguments)
	switch err {
	case flag.ErrHelp:
		help = true
	case nil:
	default:
		errorLogger.Fatalf("error parsing flags: %v", err)
	}
	// If the help flag was set, just show the help message and exit.
	if help {
		printCommandHelp(flags)
		os.Exit(0)
	}
	// verbose output?
	if verbose || configuration.VerboseOutput {
		debugLogger.SetOutput(os.Stderr)
	}
	// maxPages overwrite
	if pages != 0 {
		configuration.MaxPages = pages
	}

	args := flags.Args()
	if len(args) < 1 {
		errorLogger.Println("No parameters given!")
		printCommandHelp(flags)
		os.Exit(1)
	}
	param1 := args[0]

	debugLogger.Print("verbose flag active")
	fmt.Printf("%s called with parameter %s\n", flags.Name(), param1)

	return 0
}
