package hipchatAPI

import (
	"os"
	"testing"
)

// Dummy POST request for user input "/meetspace"
var dummyPostBodyHelper = []byte(`{"event": "room_message", "item": {"message": {"date": "2017-01-28T18:56:22.407746+00:00", "from": {"id": 123456, "links": {"self": "https://api.hipchat.com/v2/user/123456"}, "mention_name": "007", "name": "James Bond", "version": "T6JP69OQ"}, "id": "aaaaaa-bbbbbb-cccccc-eeeeee-gggggg", "mentions": [], "message": "/meetspace", "type": "message"}, "room": {"id": 333333, "is_archived": false, "links": {"members": "https://api.hipchat.com/v2/room/333333/member", "participants": "https://api.hipchat.com/v2/room/333333/participant", "self": "https://api.hipchat.com/v2/room/333333", "webhooks": "https://api.hipchat.com/v2/room/333333/webhook"}, "name": "test", "privacy": "private", "version": "RLICCCSR"}}, "oauth_client_id": "1234-5678-abcd-efgh", "webhook_id": 12345678}`)

// Dummy POST request for user input "/meetspace core
var dummyPostBodyCore = []byte(`{"event": "room_message", "item": {"message": {"date": ",2017-01-28T18:56:22.407746+00:00", "from": {"id": 123456, "links": {"self": "https://api.hipchat.com/v2/user/123456"}, "mention_name": "007", "name": "James Bond", "version": "T6JP69OQ"}, "id": "aaaaaa-bbbbbb-cccccc-eeeeee-gggggg", "mentions": [], "message": "/meetspace core", "type": "message"}, "room": {"id": 333333, "is_archived": false, "links": {"members": "https://api.hipchat.com/v2/room/333333/member", "participants": "https://api.hipchat.com/v2/room/333333/participant", "self": "https://api.hipchat.com/v2/room/333333", "webhooks": "https://api.hipchat.com/v2/room/333333/webhook"}, "name": "test", "privacy": "private", "version": "RLICCCSR"}}, "oauth_client_id": "1234-5678-abcd-efgh", "webhook_id": 12345678}`)

// init is being used here to set required env's before test execution
func init() {
	os.Setenv("MEETSPACEBOT_TEST", "true")
}

// TestParsedHipchatReq checks that the request can be wrangled into the struct
func TestParsedHipchatReq(t *testing.T) {
	parsedPost, err := ParseHipchatReq(dummyPostBodyHelper)
	if err != nil {
		t.Errorf("ParseHipchatReq() Error: %s", err)
	}

	if parsedPost.Item.Room.Name != "test" {
		t.Errorf("Expected: 'test' | Got: '%s'", parsedPost.Item.Room.Name)
	}
}

// TestCoreTeamMessage checks hipchat.NotificationRequest.Message is correct
func TestCoreRoomMessage(t *testing.T) {
	expected := `Click here to join call <a href="https://meetspaceapp.com/funtimes/core">Funtimes Core</a>`

	got, err := HipchatNotification("core", "test", "https://meetspaceapp.com/funtimes", "Funtimes", os.Getenv("MEETSPACEBOT_TEST"))
	if err != nil {
		t.Errorf("HipchatNotification() Error: %s", err)
	}

	if got != expected {
		t.Errorf("Expected: '%s' | Got: '%s'", expected, got)
	}
}

// TestHelpMessage checks that user input that doesn't match a call shows helper usage
func TestHelpMessage(t *testing.T) {
	expected := "<p><strong>Usage:</strong><br><code>/meetspace core # start core team call</code><br><code>/meetspace dev  # start dev team call</code></p>"

	got, err := HipchatNotification("", "test", "https://meetspaceapp.com/funtimes", "Funtimes", os.Getenv("MEETSPACEBOT_TEST"))
	if err != nil {
		t.Errorf("HipchatNotification() Error: %s", err)
	}

	if got != expected {
		t.Errorf("Expected: '%s' | Got: '%s'", expected, got)
	}
}
