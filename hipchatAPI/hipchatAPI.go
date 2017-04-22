package hipchatAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/davyj0nes/meetspaceBot/meetspaceAPI"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

// HipchatPostData is a struct for the data that is sent
//   When the slash command is triggered
type HipchatPostData struct {
	Event string `json:"event"`
	Item  item   `json:"item"`
}

type item struct {
	Message message `json:"message"`
	Room    struct {
		Name string `json:"name"`
	} `json:"room"`
}

type message struct {
	From    from   `json:"from"`
	Message string `json:"message"`
}

type from struct {
	MentionName string `json:"mention_name"`
	Name        string `json:"name"`
}

// Error describes the error message that is returned
type Error struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

// ParseHipchatReq simply parses the POST Req JSON into something useful
func ParseHipchatReq(data []byte) (HipchatPostData, error) {
	var parsedReq HipchatPostData
	err := json.Unmarshal(data, &parsedReq)
	if err != nil {
		return HipchatPostData{}, err
	}
	return parsedReq, nil
}

// HipchatNotification sends Room Notification Message
// The notification is chosen based on user input
func HipchatNotification(roomName, reqRoom, test string, msResponse meetspaceAPI.MeetspaceData) (string, error) {
	var notifReq *hipchat.NotificationRequest
	c := hipchat.NewClient(os.Getenv("HIPCHAT_API_TOKEN"))

	if roomName == "status" {
		notifReq = statusMessage(msResponse, roomName)
	} else if roomName != "" {
		notifReq = roomMessage(msResponse.URL, msResponse.Name, roomName)
	} else {
		notifReq = helpMessage()
	}

	// Unelegant way to test at the moment
	// Is set by environment variable
	if test == "true" {
		return notifReq.Message, nil
	}

	// This does not get covered in tests
	//  Need to find a way to mock this out
	res, err := c.Room.Notification(reqRoom, notifReq)
	// This is messy but error was returning with other implementations
	if res.StatusCode != 200 {
		if res.StatusCode != 204 {
			resMessage := Error{}
			resBody, _ := ioutil.ReadAll(res.Body)
			json.Unmarshal(resBody, &resMessage)
			return "", fmt.Errorf(" %v - \n%s\n\n%s", res.StatusCode, resMessage.Error.Message, resMessage.Error.Type)
		}
	}
	if err != nil {
		return "", err
	}

	return "", nil
}

// statusMessage is called when router sees /meetspace status
func statusMessage(msRes meetspaceAPI.MeetspaceData, slug string) *hipchat.NotificationRequest {
	meetspaceURL := fmt.Sprintf("%s/%s", strings.ToLower(msRes.URL), strings.ToLower(slug))

	// This is here as reminder that I need to get sending Cards working
	msgCard := &hipchat.Card{
		Style: "link",
		URL:   meetspaceURL,
		Title: "<p><strong>Meetspace Status</strong><br>",
		Description: hipchat.CardDescription{
			Format: "format",
			Value:  "value",
		},
	}
	msgBody := fmt.Sprintf(`%s <a href="%s"> %s</a>`, msgCard.Title, msgCard.URL, strings.Title(slug))
	return &hipchat.NotificationRequest{From: "Meetspace Bot", Message: msgBody, Color: "purple"}

	// Need to fix request for sending Card to work. More work needed
	// return &hipchat.NotificationRequest{Message: msgBody, Card: msgCard}
}

// roomMessage is called when router sees /meetspace <room name>
func roomMessage(url, team, slug string) *hipchat.NotificationRequest {
	meetspaceURL := fmt.Sprintf("%s/%s", strings.ToLower(url), strings.ToLower(team))

	// This is here as reminder that I need to get sending Cards working
	msgCard := &hipchat.Card{
		Style: "link",
		URL:   meetspaceURL,
		Title: "Click here to join call",
		Description: hipchat.CardDescription{
			Format: "format",
			Value:  "value",
		},
	}
	msgBody := fmt.Sprintf(`%s <a href="%s/%s">%s %s</a>`, msgCard.Title, msgCard.URL, strings.ToLower(slug), strings.Title(team), strings.Title(slug))
	return &hipchat.NotificationRequest{From: "Meetspace Bot", Message: msgBody, Color: "purple"}

	// Need to fix request for sending Card to work. More work needed
	// return &hipchat.NotificationRequest{Message: msgBody, Card: msgCard}
}

// helpMessage is called when any other command is given
func helpMessage() *hipchat.NotificationRequest {
	msgBody := fmt.Sprintf("<p><strong>Usage:</strong><br><code>/meetspace core # start core team call</code><br><code>/meetspace dev  # start dev team call</code></p>")

	return &hipchat.NotificationRequest{From: "Meetspace Bot", Message: msgBody, Color: "gray"}
}
