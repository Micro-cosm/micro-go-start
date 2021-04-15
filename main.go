

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)


func main() {

	log.Print("Sparking up server...")

	isLogging, _	:= strconv.ParseBool(os.Getenv("LOGS"))
	port			:= os.Getenv("REMOTE_PORT")

	if isLogging { log.Printf("defaulting to port:\t%s", port) }

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":" + port, nil);	err != nil { log.Fatal(err) }
}


func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NICKNAME")

	log.Printf("Service name:  %s", name)

	if name == "" { name = "World" }

	fprintf, err := fmt.Fprintf(w, "Hello %s!\n", name)
	if err != nil { println(fprintf)}
}