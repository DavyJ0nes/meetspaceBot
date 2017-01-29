package meetspaceAPI

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var dummyData []byte = []byte(`{"id":"1001","name":"Testy","url":"https://meetspaceapp.com/testy","rooms":[{"id":"aabb1234","name":"calltime","url":"https://meetspaceapp.com/testy/calltime","public":"\u003ctrue","participants":[{"id":"aabb1234bbaa4321","name":"jamesbond","email":"doubleoh@mod.gov","avatar-url":"https://pbs.twimg.com/profile_images/522485330771845120/gK0H2djd_400x400.jpeg"}]}]}`)

// init is being used here to set required env's before test execution
func init() {
	os.Setenv("MEETSPACE_API_TOKEN", "123456789")
}

// testServerHelper abstracts repeated code that is used to mock Meetspace API
func testServerHelper(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dummyData)

		if req.Method != "GET" {
			t.Errorf("Expected 'GET' | Got '%s'", req.Method)
		}

		if req.URL.EscapedPath() != "/i/api/v0/status" {
			t.Errorf("Expected '/i/api/v0/status' | Got '%s'", req.URL.EscapedPath())
		}

		if req.Header["Authorization"][0] != "123456789" {
			t.Errorf("Expected '123456789' | Got '%s'", req.Header["Authorization"][0])
		}
	})
}

// TestMeetSpaceRequest checks that request to API is correct
// Used to mock out Meetspace API
func TestMeetspaceRequest(t *testing.T) {
	mockHandler := testServerHelper(t)
	ts := httptest.NewServer(mockHandler)
	defer ts.Close()

	reqUrl := ts.URL
	_, err := MeetspaceCall(reqUrl, "status")
	if err != nil {
		t.Errorf("MeetspaceCall() returned error: %s", err)
	}
}

// TestMeetspaceResponse checks that byte slice from API matches dummyData
// Used to mock out Meetspace API
func TestMeetspaceResponse(t *testing.T) {
	mockHandler := testServerHelper(t)
	ts := httptest.NewServer(mockHandler)
	defer ts.Close()

	reqUrl := ts.URL
	resData, err := MeetspaceCall(reqUrl, "status")
	if err != nil {
		t.Errorf("MeetspaceCall() returned error: %s", err)
	}
	if string(resData) != string(dummyData) {
		t.Errorf("Expected: '%s' | Got: '%s'", dummyData, resData)
	}
}

// TestMeetspaceFormat checks that API response converts correctly into MeetspaceData struct
func TestMeetspaceFormat(t *testing.T) {
	mockHandler := testServerHelper(t)
	ts := httptest.NewServer(mockHandler)
	defer ts.Close()

	reqUrl := ts.URL
	resData, err := MeetspaceCall(reqUrl, "status")
	if err != nil {
		t.Errorf("MeetspaceCall() returned error: %s", err)
	}

	wrangledData, err := MeetspaceFormat(resData)
	if err != nil {
		t.Errorf("MeetspaceFormat() returned error: %s", err)
	}
	if wrangledData.Name != "Testy" {
		t.Errorf("Expected: '%s' | Got: '%s'", "Testy", wrangledData.Name)
	}
	if wrangledData.Rooms[0].Name != "calltime" {
		t.Errorf("Expected: '%s' | Got: '%s'", "calltime", wrangledData.Rooms[0].Name)
	}
	if wrangledData.Rooms[0].Participants[0].Name != "jamesbond" {
		t.Errorf("Expected: '%s' | Got: '%s'", "jamesbond", wrangledData.Rooms[0].Participants[0].Name)
	}
}
