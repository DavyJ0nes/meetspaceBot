package main

// {
// event: 'room_message',
// item: {
//     message: {
//         date: '2015-01-20T22:45:06.662545+00:00',
//         from: {
//             id: 1661743,
//             mention_name: 'Blinky',
//             name: 'Blinky the Three Eyed Fish'
//         },
//         id: '00a3eb7f-fac5-496a-8d64-a9050c712ca1',
//         mentions: [],
//         message: '/weather',
//         type: 'message'
//     },
//     room: {
//         id: 1147567,
//         name: 'The Weather Channel'
//     }
// },
// webhook_id: 578829
// }

type PostMessage struct {
	Event     string `json:"event"`
	Item      item   `json:"item"`
	WebhookID int    `json:"webhook_id"`
}

type item struct {
	Message message `json:"message"`
	Room    room    `json:"room"`
}

type message struct {
	Date     string   `json:"date"`
	From     from     `json:"from"`
	Id       string   `json:"id"`
	Mentions []string `json:"mentions"`
	Message  string   `json:"message"`
	Type     string   `json:"type"`
}

type from struct {
	Id          int    `json:"id"`
	MentionName string `json:"mention_name"`
	Name        string `json:"name"`
}

type room struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestMessage() PostMessage {
	msg := PostMessage{
		Event: "room_message",
		Item: item{
			Message: message{
				Date: "2015-01-20T22:45:06.662545+00:00",
				From: from{
					Id:          345678,
					MentionName: "Gopher",
					Name:        "Gopher",
				},
				Id:       "00a3eb7f-fac5-496a-8d64-a9050c712ca1",
				Mentions: []string{},
				Message:  "/meetspace",
				Type:     "message",
			},
			Room: room{
				Id:   987654,
				Name: "Test Room",
			},
		},
		WebhookID: 123456,
	}
	// message, err := json.Marshal(msg)
	// if err != nil {
	//   log.Fatal(err)
	// }
	return msg

}
