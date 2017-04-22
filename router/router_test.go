package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// Dummy Hipchat Request
var hcTestData = strings.NewReader(`{"event": "room_message", "item": {"message": {"date": "2017-01-28T18:56:22.407746+00:00", "from": {"id": 123456, "links": {"self": "https://api.hipchat.com/v2/user/123456"}, "mention_name": "007", "name": "James Bond", "version": "T6JP69OQ"}, "id": "aaaaaa-bbbbbb-cccccc-eeeeee-gggggg", "mentions": [], "message": "/meetspace core", "type": "message"}, "room": {"id": 333333, "is_archived": false, "links": {"members": "https://api.hipchat.com/v2/room/333333/member", "participants": "https://api.hipchat.com/v2/room/333333/participant", "self": "https://api.hipchat.com/v2/room/333333", "webhooks": "https://api.hipchat.com/v2/room/333333/webhook"}, "name": "test", "privacy": "private", "version": "RLICCCSR"}}, "oauth_client_id": "1234-5678-abcd-efgh", "webhook_id": 12345678}`)

// init is being used here to set required env's before test execution
func init() {
	os.Setenv("MEETSPACEBOT_TEST", "true")
	os.Setenv("MEETSPACE_API_HOST", "http://localhost:8080")
}

// TestHipchatHandler tests the Hipchat command route
func TestHipchatHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HipchatHandler))
	defer ts.Close()

	reqUrl := fmt.Sprintf("%s/api/v0/hipchat", ts.URL)

	res, err := http.Post(reqUrl, "", hcTestData)
	if err != nil {
		t.Errorf("Error Posting to MeetspaceHandler:, %s", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 | Got: %v", res.StatusCode)
	}
}

// TestMeetspaceData makes call to demo host running in Docker
// Not the best implementation, is more of an acceptance test
func TestMeetspaceData(t *testing.T) {
	msd, err := MeetspaceData()
	if err != nil {
		t.Errorf("Error Getting MeetspaceData():, %s", err)
	}
	if msd.Name != "Forge SP" {
		t.Errorf("Expected: 'Forge SP' | Got: '%s'", msd.Name)
	}
}
