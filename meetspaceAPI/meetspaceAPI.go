package meetspaceAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type meetspaceData struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Rooms []room `json:"rooms"`
}

type room struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Url          string        `json:"url"`
	Public       string        `json:"public"`
	Participants []participant `json:"participants"`
}

type participant struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Avatarurl string `json:"avatar-url"`
}

func MeetspaceCall(url string) ([]byte, error) {
	reqUrl := url + "/i/api/v0/status"
	res, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server didn't respond with 200 OK | Got: %s", res.Status)
	}
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func MeetspaceFormat(data []byte) error {
	return nil
}
