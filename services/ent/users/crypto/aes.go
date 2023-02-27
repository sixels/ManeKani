package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"strings"
)

func Aes256Encode(content []byte, key []byte) ([]byte, error) {
	bPlaintext := PKCS5Padding(content, aes.BlockSize)
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	iv, _ := GenerateRandomBytes(block.BlockSize())
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bPlaintext)

	ciphertextIV := append(iv, ciphertext...)

	return ciphertextIV, err
}

func Aes256Decode(cipherTextIV []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText, iv := cipherTextIV[aes.BlockSize:], cipherTextIV[:aes.BlockSize]
	out := make([]byte, len(cipherText))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(out, cipherText)

	return []byte(strings.TrimSpace(string(out))), err
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}
