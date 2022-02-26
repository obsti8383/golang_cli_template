// golang cli template

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// Configuration is the struct that gets filled by reading config.json JSON file
type Configuration struct {
	VerboseOutput bool   `json:"verbose"`
	ApiKey        string `json:"api_key"`
	MaxPages      int    `json:"max_pages"`
}

func main() {
	errorLogger := log.New(os.Stderr, "Error: ", 0)
	debugLogger := log.New(io.Discard, "", 0)

	// get configuration from config json
	var configuration Configuration
	configFile := "config.json"
	file, err := os.Open(configFile)
	if err != nil {
		// remove Fatal() in case config JSON file is optional
		errorLogger.Fatal(err.Error())
		return
	} else {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&configuration)
		if err != nil {
			errorLogger.Println(err.Error())
			return
		}
	}

	// evaluate command line flags
	var help bool
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flags.BoolVar(&help, "help", help, "Show this help message")
	flags.BoolVar(&help, "h", help, "")
	if len(os.Args) < 2 {
		printHelp(flags)
		os.Exit(2)
	}
	err = flags.Parse(os.Args[1:])
	switch err {
	case flag.ErrHelp:
		help = true
	case nil:
	default:
		errorLogger.Fatalf("error parsing flags: %v", err)
	}
	// If the help flag was set, just show the help message and exit.
	if help {
		printHelp(flags)
		os.Exit(0)
	}

	if configuration.ApiKey == "" {
		errorLogger.Println("No API key set. Please set api_key in config json.")
		printHelp(flags)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "command1":
		os.Exit(Command1(os.Args[2:], configuration, errorLogger, debugLogger))
	case "command2":
		os.Exit(Command2(os.Args[2:], configuration, errorLogger, debugLogger))
	}

	errorLogger.Println("invalid command or command missing")
	printHelp(flags)
	os.Exit(1)
}

func printHelp(flags *flag.FlagSet) {
	fmt.Fprintf(flags.Output(), "\nUsage of %s:\n", os.Args[0])
	flags.PrintDefaults()
	fmt.Printf(`
Always enter a command and at least one parameter, e.g.:
	command1 test

Possible commands are:
	command1
	command2

To configure the command, at least the api_key must be set in config.json. Example:

	{
		"verbose": false,
		"api_key": "asicj738z8fhse7h28783hiuh",
		"max_pages": 3
	}
`)
}

func printCommandHelp(flags *flag.FlagSet) {
	fmt.Fprintf(flags.Output(), "\nUsage of command %s:\n", flags.Name())
	flags.PrintDefaults()
}
