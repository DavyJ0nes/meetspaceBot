package ms_mock

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var msa MeetspaceAPI

func TestMSMock(t *testing.T) {
	mockData := []byte(`{"id":"1001","name":"Testy","url":"https://meetspaceapp.com/testy","rooms":[{"id":"aabb1234","name":"calltime","url":"https://meetspaceapp.com/testy/calltime","public":"\u003ctrue","participants":[{"id":"aabb1234bbaa4321","name":"jamesbond","email":"doubleoh@mod.gov","avatar-url":"https://pbs.twimg.com/profile_images/522485330771845120/gK0H2djd_400x400.jpeg"}]}]}`)

	err := json.Unmarshal(mockData, &msa.Data)
	if err != nil {
		panic(err)
	}
	s := httptest.NewServer(&msa)
	defer s.Close()

	resp, err := http.Get(s.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error("Expected:", 200, "Got:", resp.StatusCode)
	}

	if _, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Error(err)
	}
	// if ctype := resp.Header().Get("Content-Type"); ctype != "application/json" {
	//   t.Errorf("content type header does not match: got %v want %v",
	//     ctype, "application/json")
	// }
}
