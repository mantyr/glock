package main

import (
	"github.com/KyleBanks/glock/src/api"
	"os"
	"strings"
	"strconv"
	"github.com/KyleBanks/glock/src/glock"
	"github.com/KyleBanks/glock/src/log"
	rlog "log"
)

const (
	argPort = "-p="
	argVerbose = "-v"
)

var (
	port = 7887
	verbose = false
)

func main() {
	// Parse any command-line arguments that have been specified
	parseArgs()

	// Create the logger
	logger := log.New(verbose)

	// Create the underlying glocker
	glocker := glock.New(logger)

	// Start the API server
	defer api.New(logger, port, glocker).Run()

	// Signal that glock's ready
	logger.ForcePrintf("glock listening on %v", port)
}

// parseArgs handles the parsing of command-line arguments
func parseArgs() {
	args := os.Args[1:]

	for _, arg := range args {
		if strings.HasPrefix(arg, argPort) {
			customPort, err := strconv.Atoi(strings.Split(arg, argPort)[1])
			if err != nil {
				rlog.Fatalf("Invalid Port Specified: %v", arg)
				os.Exit(1)
			}
			port = customPort
		} else if strings.HasPrefix(arg, argVerbose) {
			verbose = true
		} else {
			rlog.Fatalf("Invalid argument: %v", arg)
			os.Exit(1)
		}
	}
}