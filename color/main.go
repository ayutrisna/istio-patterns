package main

import (
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

type PageData struct {
	Title, Color, HostName, Request string
}

func writePage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("response.html"))

	request, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	data := PageData{
		Title:    "Tugas NCP Reliable Deployment Kubernetes and Istio patterns",
		Color:    os.Getenv("COLOR"),
		HostName: hostName,
		Request:  string(request),
	}

	log.Printf("Received request %s REQ:%s\n", hostName, request)

	expiration := time.Now().Add(time.Hour)
	cookie := http.Cookie{Name: "color", Value: os.Getenv("COLOR"), Expires: expiration}
	http.SetCookie(w, &cookie)

	tmpl.Execute(w, data)

}

func main() {

	http.HandleFunc("/", writePage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
