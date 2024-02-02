package db

import (
	"backend/models"
	"fmt"

	"crypto/rand"

	"golang.org/x/crypto/nacl/box"
)

// @Summary SendMessage
// @Description SendMessage
// @ID sendMessage
// @Accept  json
// @Produce  json
// @Param username path message.SendMessage true "Message"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/message [post]
func AuthSendMessage(sender uint, reciever int, content string, attach string) (models.IResponse, int) {
	cn := Conn()

	securedMessage := string(secureMessage(content))
	m := models.Message{Sender: sender, Content: securedMessage, Reciever: reciever, Attach: attach}

	err := cn.Create(&m).Error

	if err != nil {
		return models.Response(nil, err.Error()), 400
	}

	return models.Response(
		securedMessage,
		"Message sent successfully",
	), 200
}

func secureMessage(message string) []byte {
	var firstNonce [24]byte

	rand.Read(firstNonce[:])

	myPublicKey, myPrivateKey, _ := box.GenerateKey(rand.Reader)
	yourPublicKey, yourPrivateKey, _ := box.GenerateKey(rand.Reader)

	messageToYou := []byte(message)

	encryptedForYou := box.Seal(nil, messageToYou, &firstNonce, yourPublicKey, myPrivateKey)

	decryptedForYou, _ := box.Open(nil, encryptedForYou, &firstNonce, myPublicKey, yourPrivateKey)
	fmt.Printf("%s\n", decryptedForYou)
	return encryptedForYou
}
