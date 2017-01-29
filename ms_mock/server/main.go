package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davyj0nes/meetspaceBot/ms_mock"
)

var msa ms_mock.MeetspaceAPI

func main() {
	mockData := []byte(`{"id":"1001","name":"Forge","url":"https://meetspaceapp.com/forge","rooms":[{"id":"aabb1234","name":"core","url":"https://meetspaceapp.com/forge/core","public":"\u003ctrue","participants":[{"id":"aabb1234bbaa4321","name":"jamesbond","email":"doubleoh@mod.gov","avatar-url":"https://pbs.twimg.com/profile_images/522485330771845120/gK0H2djd_400x400.jpeg"},{"id":"aabb5678bbaa4321","name":"miss moneypenny","email":"moneypenny@mi6.com","avatar-url":"https://pbs.twimg.com/profile_images/522485330771845120/gK0H2djd_400x400.jpeg"}]},{"id":"aabb1234","name":"dev","url":"https://meetspaceapp.com/forge/dev","public":"\u003ctrue","participants":[{"id":"aabb1234bbaa4321","name":"jamesbond","email":"doubleoh@mod.gov","avatar-url":"https://pbs.twimg.com/profile_images/522485330771845120/gK0H2djd_400x400.jpeg"},{"id":"aabb5678bbaa4321","name":"miss moneypenny","email":"moneypenny@mi6.com","avatar-url":"https://pbs.twimg.com/profile_images/522485330771845120/gK0H2djd_400x400.jpeg"},{"id":"aabb5678bbaa4321","name":"Q","email":"q@mi6.com","avatar-url":"https://pbs.twimg.com/profile_images/522485330771845120/gK0H2djd_400x400.jpeg"}]}]}`)

	er := json.Unmarshal(mockData, &msa.Data)
	if er != nil {
		panic(er)
	}

	http.Handle("/i/api/v0/status", &msa)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
