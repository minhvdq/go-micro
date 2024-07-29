package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "main.page.gohtml")
	})

	fmt.Println("Starting front end service on port 8081")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error no 2: ", err)
		log.Panic(err)
	}
}

//go:embed templates
var templateFS embed.FS

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"templates/base.layout.gohtml",
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		fmt.Println("Error no 1: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data struct {
		BrokerURL string
	}

	// data.BrokerURL = os.Getenv("BROKER_URL")
	data.BrokerURL = "http://localhost:8080"
	fmt.Printf("broker url is %s", os.Getenv("BROKER_URL"))

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println("Error no 3: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
