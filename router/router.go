package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/davyj0nes/meetspacebot/hipchatAPI"
	"github.com/davyj0nes/meetspacebot/meetspaceAPI"
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
	hipchatReq, hcrErr := hipchatReq(req)
	if hcrErr != nil {
		log.Fatal("Hipchat | Error parsing Body of Request: ", hcrErr)
	}

	reqRoomName := hipchatReq.Item.Room.Name
	reqMessage := hipchatReq.Item.Message.Message
	wantedCall := strings.Split(reqMessage, " ")
	meetspaceData, msdErr := MeetspaceData()
	if msdErr != nil {
		log.Fatal("Meetspace | Error calling API", msdErr)
	}

	var (
		callRoomName string
	)
	for _, room := range meetspaceData.Rooms {
		if wantedCall[len(wantedCall)-1] == room.Name {
			callRoomName = wantedCall[len(wantedCall)-1]
			break
		} else {
			callRoomName = ""
		}
	}

	_, hcnErr := hipchatAPI.HipchatNotification(callRoomName, reqRoomName, meetspaceData.Name, os.Getenv("MEETSPACEBOT_TEST"))
	if hcnErr != nil {
		log.Fatal("Hipchat | Error Sending Request: ", hcnErr)
	}
	w.Header().Set("Content-Type", "application/json")
}

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

func MeetspaceData() (meetspaceAPI.MeetspaceData, error) {
	apiUrl := os.Getenv("MEETSPACE_API_HOST")
	apiEndpoint := "status"
	apiReq, err := meetspaceAPI.MeetspaceCall(apiUrl, apiEndpoint)
	if err != nil {
		return meetspaceAPI.MeetspaceData{}, err
	}
	return meetspaceAPI.MeetspaceFormat(apiReq)
}
