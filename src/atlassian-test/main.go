package main

import (
	_ "atlassian-test/emoicons"
	_ "atlassian-test/links"
	_ "atlassian-test/mentions"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

func main() {
	
	// !Warning: Setting log level, default.
	
	loglevel, err := log.ParseLevel(getEnv("LOGLEVEL", "warn"))
	if err == nil {
		log.SetLevel(loglevel)
	}
	
	log.Printf("Log level: %s", log.GetLevel())
	// Starting HTTP server with one handler
	http.HandleFunc("/filter", filterHandler)
	log.Fatal(http.ListenAndServe(getEnv("LISTEN", ":8080"), nil))
}
