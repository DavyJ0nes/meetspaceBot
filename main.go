package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// {
// "color": "green",
// "message": "It's going to be sunny tomorrow! (yey)",
// "notify": false,
// "message_format": "text"
// }
func testHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	reqBody, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(reqBody))
	fmt.Fprintf(w, "Hey")
}

func main() {
	http.HandleFunc("/", testHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
