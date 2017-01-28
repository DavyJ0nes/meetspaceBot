package router

import (
	"log"
	"net/http"

	"github.com/davyj0nes/meetspacebot/hipchatAPI"
)

func requestLogger(req *http.Request) {
	log.Printf("%s | %s || %s => %s || %s", req.Method, req.URL.Path, req.RemoteAddr, req.Host, req.Header.Get("User-Agent"))
}

func Router() {
	http.HandleFunc("/api/v0/hipchat", HipchatHandler)
	http.HandleFunc("/api/v0/slack", SlackHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func HipchatHandler(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)
	// reqBody, _ := ioutil.ReadAll(req.Body)
	// log.Printf("Body: '%s'", string(reqBody))
	w.Header().Set("Content-Type", "application/json")
	err := hipchatAPI.HipchatNotification()
	if err != nil {
		log.Fatal(err)
	}
}

func SlackHandler(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hey"))
}
