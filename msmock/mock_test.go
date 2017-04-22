package msmock

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var msa MeetspaceAPI

func TestMSMock(t *testing.T) {
	mockData := []byte(`{"name":"Testy","url":"https://meetspaceapp.com/test","rooms":[{"name":"calltime","url":"https://meetspaceapp.com/test/calltime","public":false,"participants":[]},{"name":"Other Room","url":"https://meetspaceapp.com/test/other","public":true,"participants":[{"name":"James Bond","email":"bond@doubleoh.com","avatar-url":"http://vignette2.wikia.nocookie.net/jamesbond/images/d/de/James_Bond_(Roger_Moore)_-_Profile.jpg"}]}]}`)

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
