package db

import (
	"backend/models"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func GetMessages(chatId uint, cursor int) (models.IResponse, int) {
	cn := Conn()

	var m []models.Message

	cn = cn.
		Limit(40).
		Order("id DESC").
		Where("chat_id = ?", chatId)
	if cursor != 0 {
		cn = cn.Where("id < ?", cursor)
	}
	err := cn.Find(&m).Error

	for i := 0; i < len(m); i++ {
		m[i].Content, err = GetAESDecrypted(m[i].Content)
		if err != nil {
			return models.Response(nil, err.Error()), 401
		}
	}

	if err != nil {
		return models.Response(nil, err.Error()), 400
	}
	return models.Response(m, "success"), 200
}

func SendMessage(sender uint, chatId uint, content string, attach string) (models.IResponse, int) {
	cn := Conn()

	var encrypted string
	var err error
	if encrypted, err = GetAESEncrypted(content); err != nil {
		return models.Response(nil, err.Error()), 400
	}

	m := models.Message{Sender: sender, Content: encrypted, ChatId: chatId, Attach: attach}

	err = cn.Create(&m).Error

	if err != nil {
		return models.Response(nil, err.Error()), 400
	}

	return models.Response(
		map[string]interface{}{
			"message": content,
			"attach":  attach,
		},
		"Message sent successfully",
	), 200
}

func GetAESEncrypted(plaintext string) (string, error) {
	key := "my32digitkey12345678901234567890"
	iv := "my16digitIvKey12"

	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.StdEncoding.EncodeToString(ciphertext)

	return str, nil
}

func GetAESDecrypted(encrypted string) (string, error) {
	key := "my32digitkey12345678901234567890"
	iv := "my16digitIvKey12"

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = pkcs5Unpadding(ciphertext)

	return string(ciphertext), nil
}

func pkcs5Unpadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}
