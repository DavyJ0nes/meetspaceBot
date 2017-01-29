package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var hcTestData = strings.NewReader(`{"event": "room_message", "item": {"message": {"date": "2017-01-28T18:56:22.407746+00:00", "from": {"id": 123456, "links": {"self": "https://api.hipchat.com/v2/user/123456"}, "mention_name": "007", "name": "James Bond", "version": "T6JP69OQ"}, "id": "aaaaaa-bbbbbb-cccccc-eeeeee-gggggg", "mentions": [], "message": "/meetspace", "type": "message"}, "room": {"id": 333333, "is_archived": false, "links": {"members": "https://api.hipchat.com/v2/room/333333/member", "participants": "https://api.hipchat.com/v2/room/333333/participant", "self": "https://api.hipchat.com/v2/room/333333", "webhooks": "https://api.hipchat.com/v2/room/333333/webhook"}, "name": "test", "privacy": "private", "version": "RLICCCSR"}}, "oauth_client_id": "1234-5678-abcd-efgh", "webhook_id": 12345678}`)

// init is being used here to set required env's before test execution
func init() {
	os.Setenv("MEETSPACEBOT_TEST", "true")
}

// func TestRouter(t *testing.T) {
//   req, err := http.NewRequest("GET", "/api/v0/hipchat", nil)
//   if err != nil {
//     t.Fatal(err)
//   }

//   rr := httptest.NewRecorder()
//   handler := http.HandlerFunc(HipchatHandler)

//   handler.ServeHTTP(rr, req)
//   if status := rr.Code; status != http.StatusOK {
//     t.Errorf("handler returned wrong status code: got %v want %v",
//       status, http.StatusOK)
//   }
// }

func TestHipchatHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HipchatHandler))
	defer ts.Close()

	reqUrl := fmt.Sprintf("%s/api/v0/hipchat", ts.URL)

	res, err := http.Post(reqUrl, "", hcTestData)
	if err != nil {
		t.Errorf("Error Posting to MeetspaceHandler:, %s", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 | Got: %s", res.StatusCode)
	}
}
