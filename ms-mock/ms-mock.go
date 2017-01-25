package mock

import (
	"encoding/json"
	"net/http"
)

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
type meetspaceData struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Rooms []room `json:"rooms"`
}

type MeetspaceAPI struct {
	Data meetspaceData
}

func (api *MeetspaceAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api.Data)
}
