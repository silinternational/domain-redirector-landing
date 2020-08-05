package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	additionalMessage template.HTML
	newHost           string
	moreInfoURL       string
	port              string
	redirectEndDate   string
)

type pageData struct {
	OldHost           string
	NewHost           string
	NewURL            string
	MoreInfoURL       string
	RedirectEndDate   string
	AdditionalMessage template.HTML
}

func loadConfig() {
	newHost = os.Getenv("NEW_HOST")
	if newHost == "" {
		log.Println("NEW_HOST env var is required")
		os.Exit(1)
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// optional vars
	moreInfoURL = os.Getenv("MORE_INFO_URL")
	redirectEndDate = os.Getenv("REDIRECT_END_DATE")
	additionalMessage = template.HTML(os.Getenv("ADDITIONAL_MESSAGE"))
}

func main() {
	loadConfig()
	http.HandleFunc("/", serveTemplate)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	data := getPageData(r)

	tmpl, err := template.New("index").Parse(pageTemplate)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func getPageData(r *http.Request) pageData {
	query := ""
	if len(r.URL.Query()) > 0 {
		query = "?" + r.URL.Query().Encode()
	}

	return pageData{
		OldHost:           r.Host,
		NewHost:           newHost,
		NewURL:            fmt.Sprintf("https://%s%s%s", newHost, r.URL.Path, query),
		MoreInfoURL:       moreInfoURL,
		RedirectEndDate:   redirectEndDate,
		AdditionalMessage: additionalMessage,
	}
}
