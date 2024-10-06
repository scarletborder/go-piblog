package main

import (
	"flag"
	"fmt"
	"os"
)

var Version = "0.0.1a"

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: piblog <comment> [options]")
		fmt.Println("\nCommands:")
		fmt.Println("	upload	Manage upload")
		fmt.Println("	stat	Inspect statistic information")
		fmt.Println("\nFlags:")
		fmt.Println("	-h		Get help")
		fmt.Println("	-v		Get version info")

	}

	if len(os.Args) <= 2 {
		if len(os.Args) == 1 {
			flag.Usage()
		} else {
			switch os.Args[1] {
			case "-h":
				flag.Usage()
			case "-v":
				fmt.Printf("piblog version  %s\n", Version)
			default:
				flag.Usage()
				os.Exit(1)
			}
		}
		os.Exit(0)
	}

	switch os.Args[1] {
	case "upload":
		handleUploadCommand(os.Args[2:])
	case "Stat":
		handleStatCommand(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		flag.Usage()
		os.Exit(1)
	}
}
