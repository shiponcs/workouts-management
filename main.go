package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/shiponcs/femProject/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "The backend server port")
	flag.Parse()
	fmt.Println(port)
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Println("The app is running at ", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
}
