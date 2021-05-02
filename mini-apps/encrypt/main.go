package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// Encrypt takes in a key and plaintext values, returns a encrypted
func Encrypt(key, plaintext string) (string, error) {

	block, err := newCipherBlock(key)

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x\n", ciphertext), nil
}

// Decrypt will take a key and cipherHex and return a decrypted string
func Decrypt(key, cipherHex string) (string, error) {

	block, err := newCipherBlock(key)
	if err != nil {
		panic(err)
	}

	ciphertext, err := hex.DecodeString(cipherHex)

	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: cipher too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()

	fmt.Fprint(hasher, key)
	ciperKey := hasher.Sum(nil)
	block, err := aes.NewCipher(ciperKey)

	if err != nil {
		return nil, err
	}

	return block, nil
}
