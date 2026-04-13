package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var version string = "0.0.0"

func main() {
	v := flag.Bool("v", false, "print version")

	flag.Parse()

	if *v {
		fmt.Println(version)
		return
	}

	log.Println("Starting...")

	port := getEnvOrDefault("PORT", "8080")

	addr := fmt.Sprintf(":%v", port)

	mux := http.NewServeMux()

	addHandlers(mux)

	log.Printf("Listening on http://localhost:%v", port)

	err := http.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatalln(err)
	}
}

func addHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", handleIndex)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	now := time.Now()

	w.Write(getResponse(now))
}

func getResponse(now time.Time) []byte {
	res := fmt.Sprintf("Hello World! The time is %v", now)
	return []byte(res)
}

func getEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
