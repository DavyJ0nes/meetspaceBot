package meetspaceAPI

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMeetspaceRequest(t *testing.T) {
	mockStatus := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)

		if req.Method != "GET" {
			t.Errorf("Expected 'GET' | Got '%s'", req.Method)
		}

		if req.URL.EscapedPath() != "/i/api/v0/status" {
			t.Errorf("Expected '/i/api/v0/status' | Got '%s'", req.URL.EscapedPath())
		}
	})

	ts := httptest.NewServer(mockStatus)
	defer ts.Close()

	reqUrl := ts.URL
	_, err := MeetspaceCall(reqUrl)
	if err != nil {
		t.Errorf("MeetspaceCall() returned error: %s", err)
	}
}

func TestMeetspaceResponse(t *testing.T) {
	dummyData := []byte(`{"id":"1001","name":"Testy","url":"https://meetspaceapp.com/testy","rooms":[{"id":"aabb1234","name":"calltime","url":"https://meetspaceapp.com/testy/calltime","public":"\u003ctrue","participants":[{"id":"aabb1234bbaa4321","name":"jamesbond","email":"doubleoh@mod.gov","avatar-url":"https://pbs.twimg.com/profile_images/522485330771845120/gK0H2djd_400x400.jpeg"}]}]}`)
	mockStatus := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dummyData)
	})

	ts := httptest.NewServer(mockStatus)
	defer ts.Close()

	reqUrl := ts.URL
	data, err := MeetspaceCall(reqUrl)
	if err != nil {
		t.Errorf("MeetspaceCall() returned error: %s", err)
	}
	if string(data) != string(dummyData) {
		t.Errorf("Expected: '%s' | Got: '%s'", dummyData, data)
	}
}
