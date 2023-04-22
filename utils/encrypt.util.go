package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"wimb-backend/config"

	"golang.org/x/crypto/pbkdf2"
)

func GetEncryptStream(iv []byte) (cipher.Stream, error) {

	key := pbkdf2.Key([]byte(config.Env.PASSWORD), []byte(config.Env.HASH_SALT), 1000, 32, sha256.New)

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)

	return stream, nil

}

func Encrypt(value string) (string, error) {

	iv := make([]byte, aes.BlockSize)

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream, err := GetEncryptStream(iv)

	if err != nil {
		return "", err
	}

	cipherText := make([]byte, len(value))
	stream.XORKeyStream(cipherText, []byte(value))

	encrypted := fmt.Sprintf("%x%x", iv, cipherText)

	return encrypted, nil
}

func Decrypt(value string) (string, error) {

	encryptedBytes, err := hex.DecodeString(value)

	if err != nil {
		return "", err
	}

	ivBytes := encryptedBytes[:aes.BlockSize]
	chiperTextBytes := encryptedBytes[aes.BlockSize:]

	stream, err := GetEncryptStream(ivBytes)

	if err != nil {
		return "", err
	}

	plaintext := make([]byte, len(chiperTextBytes))
	stream.XORKeyStream(plaintext, chiperTextBytes)

	return string(plaintext), nil
}
