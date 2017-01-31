package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davyj0nes/meetspaceBot/ms_mock"
)

var msa ms_mock.MeetspaceAPI

func main() {
	mockData := []byte(`{"name":"Forge SP","url":"https://meetspaceapp.com/forge","rooms":[{"name":"Forge Core Team","url":"https://meetspaceapp.com/forge/core","public":false,"participants":[]},{"name":"Exposure Dev Team","url":"https://meetspaceapp.com/forge/dev","public":true,"participants":[]},{"name":"Robs Room","url":"https://meetspaceapp.com/forge/robsroom","public":true,"participants":[]},{"name":"Davy Open Room","url":"https://meetspaceapp.com/forge/davy","public":true,"participants":[]}]}`)

	er := json.Unmarshal(mockData, &msa.Data)
	if er != nil {
		panic(er)
	}

	http.Handle("/i/api/v0/status", &msa)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
