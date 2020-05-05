package main

import (
	u "sf-logger/src/utils"
)

func main() {
	// Strting the Application
	u.Info.Println("Starting..")

	// Execution od Code
	u.Debug.Println("Running Process 1")
	u.Debug.Println("Running Process 2")

	u.Error.Println("An Error Occured: Error abcd")

	u.Debug.Println("Running Process i")
	u.Error.Println("An Error Occured: Error efgh")
	u.Debug.Println("Running Process n")

	// Ending the Application
	u.Info.Println("Execution Completed..")
}
