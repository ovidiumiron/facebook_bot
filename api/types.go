package api

import "time"

type Message struct {
	Mid  string `json:"mid"`
	Seq  int64  `json:"seq"`
	Test string `json:"text"`
}

type Sender struct {
	Id string `json:"id"`
}

type Recipient struct {
	Id string `json:"id"`
}

type Messages struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp time.Time `json: "timestamp"`
	Message   Message   `json:"message"`
}

type WebHookEvent struct {
	Object string     `json:"object"`
	Entry  []Messages `json:"entry"`
}
