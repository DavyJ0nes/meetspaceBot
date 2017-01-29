package meetspaceAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// MeetspaceData defines API response from meetspace
type MeetspaceData struct {
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

// MeetspaceCall makes API call
// Am hardcoding the API version at v0 to stop breaking changes in future
//  major version releases
func MeetspaceCall(url string, endpoint string) ([]byte, error) {
	client := &http.Client{}
	reqUrl := url + "/i/api/v0/" + endpoint
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}

	// Using Env to abstract keys to from app
	token := os.Getenv("MEETSPACE_API_TOKEN")
	req.Header.Set("Authorization", token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server Replied with: %s", res.Status)
	}
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MeetspaceFormat parses API JSON response into MeetspaceData struct
//  for use in other parts of the app
func MeetspaceFormat(data []byte) (MeetspaceData, error) {
	var formattedData MeetspaceData
	err := json.Unmarshal(data, &formattedData)
	if err != nil {
		return MeetspaceData{}, err
	}
	return formattedData, nil
}
