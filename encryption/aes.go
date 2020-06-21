package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func EncryptAES(key string, srcText []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil, err
	}

	padding := block.BlockSize() - len(srcText)%block.BlockSize()
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	srcText = append(srcText, padtext...)

	blockModel := cipher.NewCBCEncrypter(block, []byte(key))

	ciphertext := make([]byte, len(srcText))

	blockModel.CryptBlocks(ciphertext, srcText)
	return ciphertext, nil
}

func Decrypt(key string, srcText []byte) ([]byte, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)

	if err != nil {
		return nil, err
	}

	blockModel := cipher.NewCBCDecrypter(block, keyBytes)
	plantText := make([]byte, len(srcText))
	blockModel.CryptBlocks(plantText, srcText)
	unpadding := int(plantText[len(plantText)-1])

	return plantText[:(len(plantText) - unpadding)], nil
}
