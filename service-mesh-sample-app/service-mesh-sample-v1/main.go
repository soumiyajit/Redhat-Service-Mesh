package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/helloworld", helloworldHandler)
	http.HandleFunc("/api/v1/helloiworld", helloworldHandler)

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Redhat Service Mesh by Soumiyajit")
}

type helloworld struct {
	Message string
	Version string
}

func helloworldHandler(w http.ResponseWriter, r *http.Request) {

	m := helloworld{"Service Mesh Sample App by Soumiyajit.", "v1"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
