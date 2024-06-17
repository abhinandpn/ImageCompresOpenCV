package model

// webhook responce body

type WebhookReqBody struct {
	Message Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	ID int64 `json:"id"`
}

// type WebhookReqBody struct {
// 	Message struct {
// 		Text string `json:"text"`
// 		Chat struct {
// 			ID int64 `json:"id"`
// 		} `json:"chat"`
// 	} `json:"message"`
// }
