package Actions

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func EncryptAES(input, key, iv string, blockSize int ) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bInput := PKCS5Padding([]byte(input), blockSize, len(input))
	block, _ := aes.NewCipher(bKey)
	output := make([]byte, len(bInput))
	
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(output, bInput)
	return string(output)
}


func DecryptAES(input, key, iv interface{}, blockSize interface{} ) string {

	text := []byte(input.(string))
	bKey := []byte(key.(string))
	bIV := []byte(iv.(string))
	
	block, _ := aes.NewCipher(bKey)
	if len(text) < aes.BlockSize {
		panic( "Ciphertext too short")
	}
	
	decrypted := make([]byte, len(text))
	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks(decrypted, text)
	
	
	return string(PKCS5UnPadding(decrypted))
}



func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}