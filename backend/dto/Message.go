package dto

type Message struct {
	Reciever int    `json:"reciever"`
	Content  string `json:"content"`
	Attach   string `json:"attach"`
}
