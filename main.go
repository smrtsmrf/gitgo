package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

// flags
var (
	user string
)

func main() {
	flag.Parse() // parse flags

	// if user don't supply flags, print usage
	if flag.NFlag() == 0 {
    printUsage()
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func printUsage() {
  fmt.Printf("Usage: %s [options]\n", os.Args[0])
  fmt.Println("Options:")
  flag.PrintDefaults()
  os.Exit(1)
}