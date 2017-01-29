package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/davyj0nes/meetspacebot/hipchatAPI"
)

func requestLogger(req *http.Request) {
	log.Printf("%s | %s || %s => %s || %s", req.Method, req.URL.Path, req.RemoteAddr, req.Host, req.Header.Get("User-Agent"))
}

func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/hipchat", HipchatHandler)
	// mux.HandleFunc("/api/v0/slack", SlackHandler)
	log.Fatal(http.ListenAndServe(":8081", mux))
	return mux
}

func HipchatHandler(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal("Couldn't read Body of Request", err)
	}

	parsed, err := hipchatAPI.ParseHipchatReq(reqBody)
	if err != nil {
		log.Fatal("Hipchat | Error parsing Body of Request: ", err)
	}

	_, er := hipchatAPI.HipchatNotification(parsed, os.Getenv("MEETSPACEBOT_TEST"))
	if er != nil {
		log.Fatal("Hipchat | Error Sending Request: ", er)
	}
	w.Header().Set("Content-Type", "application/json")
}
