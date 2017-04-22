package msmock

import (
	"encoding/json"
	"net/http"
)

type room struct {
	Name         string        `json:"name"`
	URL          string        `json:"url"`
	Public       bool          `json:"public"`
	Participants []participant `json:"participants"`
}

type participant struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Avatarurl string `json:"avatar-url"`
}
type meetspaceData struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Rooms []room `json:"rooms"`
}

// MeetspaceAPI is a mocked represenentation of the intended API schema
type MeetspaceAPI struct {
	Data meetspaceData
}

func (api *MeetspaceAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api.Data)
}
