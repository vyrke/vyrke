package main

import (
	"fmt"
	"flag"
)

func help() {
	fmt.Println("Vyrke CLI: Deploy by yesterday")
	fmt.Println("\nUsage:\n  vyrke [options] <command>")
	fmt.Println("\n  Example:\n    vyrke deploy")
	fmt.Println("\nCommands:")
	fmt.Println("  deploy	deploy to the vyrke platform (default)")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
}