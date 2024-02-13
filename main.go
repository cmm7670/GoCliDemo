package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	name := flag.String("name", "user", "The name that will be used in messages.")

	command := flag.String("command", "", "The command to execute.")

	flag.Parse()

	if *command == "" && flag.NArg() == 0 {
		fmt.Printf("Please provide a command with the -command flag or as an argument, %s.\n", *name)
		flag.Usage()
		return
	}

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

	if flag.Arg(0) == "time" {
		fmt.Printf("Current time as requested by %s: %s\n", *name, time.Now().Format("15:04:05"))
	} else {
		fmt.Printf("The time was not requested by %s.\n", *name)
	}

	fmt.Printf("Message for %s: End of program.\n", *name)
}
