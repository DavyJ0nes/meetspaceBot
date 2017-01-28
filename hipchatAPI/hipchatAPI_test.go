package hipchatAPI

import "testing"

var dummyPostBody = []byte(`{"event": "room_message", "item": {"message": {"date": "2017-01-28T18:56:22.407746+00:00", "from": {"id": 123456, "links": {"self": "https://api.hipchat.com/v2/user/123456"}, "mention_name": "007", "name": "James Bond", "version": "T6JP69OQ"}, "id": "aaaaaa-bbbbbb-cccccc-eeeeee-gggggg", "mentions": [], "message": "/meetspace", "type": "message"}, "room": {"id": 333333, "is_archived": false, "links": {"members": "https://api.hipchat.com/v2/room/333333/member", "participants": "https://api.hipchat.com/v2/room/333333/participant", "self": "https://api.hipchat.com/v2/room/333333", "webhooks": "https://api.hipchat.com/v2/room/333333/webhook"}, "name": "test", "privacy": "private", "version": "RLICCCSR"}}, "oauth_client_id": "1234-5678-abcd-efgh", "webhook_id": 12345678}`)

func TestParsedHipchatReq(t *testing.T) {
	parsedPost, err := ParseHipchatReq(dummyPostBody)
	if err != nil {
		t.Errorf("ParseHipchatReq() Error: %s", err)
	}

	// At the moment I only care about the Room Name
	if parsedPost.Item.Room.Name != "test" {
		t.Errorf("Expected: 'test' | Got: '%s'", parsedPost.Item.Room.Name)
	}
}
