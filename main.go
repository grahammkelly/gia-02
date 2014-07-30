// gia-02
package main

import (
	"log"
	"os"

	_ "./matchers"
	_ "./search"
)

// init is executed prior to main()
func init() {
	//Change the device for logging
	log.SetOutput(os.Stdout)
}

//Main entry point for the application
func main() {
	log.Println("Starting search!")

	//Perform the logic (search for the term)
	//search.Run("president")

	log.Println("Finished search")
}
