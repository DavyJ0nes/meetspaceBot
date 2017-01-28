package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHipchatHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HipchatHandler))
	defer ts.Close()
	reqUrl := ts.URL
	res, err := http.Post(reqUrl, "", nil)
	if err != nil {
		t.Errorf("Error Posting to MeetspaceHandler:, %s", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestSlackHandler(t *testing.T) {
}
