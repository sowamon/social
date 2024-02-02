package dto

type Message struct {
	Reciever int    `json:"reciever" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Attach   string `json:"attach"`
}
