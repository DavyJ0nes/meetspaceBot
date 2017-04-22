package meetspaceAPI

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var dummyData = []byte(`{"name":"Testy","url":"https://meetspaceapp.com/test","rooms":[{"name":"calltime","url":"https://meetspaceapp.com/test/calltime","public":false,"participants":[]},{"name":"Other Room","url":"https://meetspaceapp.com/test/other","public":true,"participants":[{"name":"James Bond","email":"bond@doubleoh.com","avatar-url":"http://vignette2.wikia.nocookie.net/jamesbond/images/d/de/James_Bond_(Roger_Moore)_-_Profile.jpg"}]}]}`)

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

	reqURL := ts.URL
	_, err := MeetspaceCall(reqURL, "status")
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

	reqURL := ts.URL
	resData, err := MeetspaceCall(reqURL, "status")
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

	reqURL := ts.URL
	resData, err := MeetspaceCall(reqURL, "status")
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
	if wrangledData.Rooms[1].Participants[0].Name != "James Bond" {
		t.Errorf("Expected: '%s' | Got: '%s'", "James Bond", wrangledData.Rooms[1].Participants[0].Name)
	}
}
