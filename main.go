package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var version string = "0.0.0"
var dbConn string
var started = time.Now()

func main() {
	v := flag.Bool("v", false, "print version")

	flag.Parse()

	if *v {
		fmt.Println(version)
		return
	}

	log.Println("Starting...")

	godotenv.Load()

	dbConn = getEnvOrDefault("DB_CONN", "")
	port := getEnvOrDefault("PORT", "8080")

	addr := fmt.Sprintf(":%v", port)

	mux := http.NewServeMux()

	addHandlers(mux)

	handler := loggingMiddleware(mux)

	log.Printf("Listening on http://localhost:%v", port)

	err := http.ListenAndServe(addr, handler)

	if err != nil {
		log.Fatalln(err)
	}
}

func addHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/about", handleAbout)
	mux.HandleFunc("/", handleIndex)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%s)", r.Method, r.URL.Path, time.Since(start))
	})
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	now := time.Now()

	w.Write(getResponse(now))
}

func getResponse(now time.Time) []byte {
	dbStatus := tryGetDbStatus()
	res := fmt.Sprintf("Hello World! The time is %v. %v.", now, dbStatus)
	return []byte(res)
}

func tryGetDbStatus() string {
	status, err := getDbStatus()

	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}

	return status
}

func getDbStatus() (string, error) {
	if dbConn == "" {
		return "There is no DB_CONN set", nil
	}

	db, err := sql.Open("postgres", dbConn)

	if err != nil {
		return "", err
	}

	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)

	if err != nil {
		return "", err
	}

	status := fmt.Sprintf("Connected to %v", version)

	return status, nil
}

func getEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	res := fmt.Sprintf("Hello World Complete\nversion: %v\nstarted: %v", version, started)
	w.Write([]byte(res))
}
