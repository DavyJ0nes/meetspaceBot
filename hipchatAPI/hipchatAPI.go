package hipchatAPI

import (
	"encoding/json"
	"os"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

// example POST message
// {
// "event": "room_message",
// "item": {
//   "message": {
//     "date": "2017-01-28T18:56:22.407746+00:00",
//     "from": {
//       "id": 3709716,
//       "links": {
//         "self": "https://api.hipchat.com/v2/user/3709716"
//       },
//       "mention_name": "davy",
//       "name": "Davy Jones",
//       "version": "T6JP69OQ"
//     },
//     "id": "d6ebec9f-1fa5-4e65-8f2b-54745150b718",
//     "mentions": [],
//     "message": "/meetspace",
//     "type": "message"
//   },
//   "room": {
//     "id": 3143303,
//     "is_archived": false,
//     "links": {
//       "members": "https://api.hipchat.com/v2/room/3143303/member",
//       "participants": "https://api.hipchat.com/v2/room/3143303/participant",
//       "self": "https://api.hipchat.com/v2/room/3143303",
//       "webhooks": "https://api.hipchat.com/v2/room/3143303/webhook"
//     },
//     "name": "hw",
//     "privacy": "private",
//     "version": "RLICVYSR"
//   }
// },
// "oauth_client_id": "1b9f9174-a01f-40d1-9e4b-415e405c5b5b",
// "webhook_id": 16080298
// }

// HipchatPostData is a struct for the data that is sent
//   When the slash command is triggered
// Only need Room Name from the POST request
type HipchatPostData struct {
	Event string `json:"event"`
	Item  item   `json:"item"`
}

type item struct {
	Room room `json:"room"`
}

type room struct {
	Name string `json:"name"`
}

// ParsedHipchatReq simply parses the POST Req JSON into something useful
func ParseHipchatReq(data []byte) (HipchatPostData, error) {
	var parsedReq HipchatPostData
	err := json.Unmarshal(data, &parsedReq)
	if err != nil {
		return HipchatPostData{}, err
	}
	return parsedReq, nil
}

// HipchatNotification sends Room Notification Message
func HipchatNotification() error {
	hc_auth := os.Getenv("HIPCHAT_API_TOKEN")
	c := hipchat.NewClient(hc_auth)
	notifMsg := &hipchat.NotificationRequest{Message: "Testing"}
	_, err := c.Room.Notification("hw", notifMsg)
	if err != nil {
		return err
	}
	return nil
}
