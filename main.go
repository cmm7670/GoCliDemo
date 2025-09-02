package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	// Define two flags that the CLI will accept.
	// The name flag has a default value of "user" and the command flag has no default value.
	// The third parameter is the help message for the flag.
	name := flag.String("name", "user", "The name that will be used in messages.")
	command := flag.String("command", "", "The command to execute.")

	// Parse the command line into defined flags
	flag.Parse()

	// If the command flag was not provided and there are no arguments,
	// print a message and the usage info.
	if *command == "" && flag.NArg() == 0 {
		fmt.Printf("Please provide a command with the -command flag or as an argument, %s.\n", *name)
		flag.Usage()
		return
	}

	// Check the value of the command flag and execute the appropriate code.
	switch *command {
	case "random":
		fmt.Printf("Generated a random number between 1 and 100 for %s: %d\n", *name, rand.Intn(100))
	case "list":
		files, _ := os.Open(".")
		defer files.Close()

		fileInfo, _ := files.Readdir(-1)
		for _, file := range fileInfo {
			fmt.Println(file.Name())
		}
	}

	// Check the value of the first argument and print the time if requested.
	if flag.Arg(0) == "time" {
		fmt.Printf("Current time as requested by %s: %s\n", *name, time.Now().Format("15:04:05"))
	} else {
		fmt.Printf("The time was not requested by %s.\n", *name)
	}

	fmt.Printf("Message for %s: End of program.\n", *name)
}
