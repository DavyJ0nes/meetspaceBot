package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/davyj0nes/meetspaceBot/hipchatAPI"
	"github.com/davyj0nes/meetspaceBot/meetspaceAPI"
)

// requestLogger logs request information in standard way
func requestLogger(req *http.Request) {
	log.Printf(">> %s | %s || %s => %s || %s", req.Method, req.URL.Path, req.RemoteAddr, req.Host, req.Header.Get("User-Agent"))
}

// Router is main mux wrangler. Keeps main() clean
func Router() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v0/hipchat", HipchatHandler)
	// mux.HandleFunc("/api/v0/slack", SlackHandler)
	log.Fatal(http.ListenAndServe(":8081", mux))
}

// HipchatHandler is the main function for dealing with Hipchat Requests
func HipchatHandler(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)
	hipchatReq, hcrErr := hipchatReq(req)
	if hcrErr != nil {
		log.Fatal("Hipchat | Error parsing Body of Request: ", hcrErr)
	}

	var callRoomName string
	reqRoomName := hipchatReq.Item.Room.Name
	reqMessage := hipchatReq.Item.Message.Message
	wantedCall := strings.Split(reqMessage, " ")
	meetspaceData, msdErr := MeetspaceData()
	if msdErr != nil {
		log.Fatal("Meetspace | Error calling API", msdErr)
	}

	for _, room := range meetspaceData.Rooms {
		if len(wantedCall) < 1 {
			callRoomName = ""
			break
		}
		if strings.Contains(room.URL, wantedCall[len(wantedCall)-1]) {
			callRoomName = wantedCall[len(wantedCall)-1]
			break
		} else {
			callRoomName = ""
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_, hcnErr := hipchatAPI.HipchatNotification(callRoomName, reqRoomName, os.Getenv("MEETSPACEBOT_TEST"), meetspaceData)
	if hcnErr != nil {
		log.Fatal("Hipchat | Error Sending Request: ", hcnErr)
	}
}

// hipchatReq returns formatted Data from Hipchat POST request
func hipchatReq(req *http.Request) (hipchatAPI.HipchatPostData, error) {
	var hcpd hipchatAPI.HipchatPostData
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return hcpd, err
	}

	parsed, err := hipchatAPI.ParseHipchatReq(reqBody)
	if err != nil {
		return hcpd, err
	}
	return parsed, nil
}

// MeetspaceData calls the meetspace API and returns formatted Data
func MeetspaceData() (meetspaceAPI.MeetspaceData, error) {
	apiURL := os.Getenv("MEETSPACE_API_HOST")
	apiEndpoint := "status"
	apiReq, err := meetspaceAPI.MeetspaceCall(apiURL, apiEndpoint)
	if err != nil {
		return meetspaceAPI.MeetspaceData{}, err
	}
	return meetspaceAPI.MeetspaceFormat(apiReq)
}
