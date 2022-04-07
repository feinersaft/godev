// This is the package, which groups all functions as well
//  as files in the same directory.
package main

// Import dependency modules
import (
	"fmt"
	"log"

	"greetings"
)

// A main function executes by default when you run the main package
func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

	// Request a greeting message.
    message, err := greetings.Hellos(names)
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

	// If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}