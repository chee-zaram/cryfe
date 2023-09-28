package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// readInputFile calls os.ReadFile to read a file, and returns the plain text
// and any error that may have occured.
func readInputFile(inputFile string) ([]byte, error) {
	return os.ReadFile(inputFile)
}

// createAES_CipherBlock takes in a key and returns a new aes cipher block and
// any error which may have occured.
func createAES_CipherBlock(key []byte) (cipher.Block, error) {
	if key == nil {
		return nil, fmt.Errorf("key cannot be nil")
	}

	return aes.NewCipher(key)
}

// readRandBytesToBuf takes a slice of bytes and reads randome bytes into it
// from rand.Reader.
func readRandBytesToBuf(buf []byte) (int, error) {
	return io.ReadFull(rand.Reader, buf)
}

// EncryptFile encrypts a file using AES algorithm and returns an error if any.
func EncryptFile(inputFile, outputFile string, key []byte) error {
	plainText, err := readInputFile(inputFile)
	if err != nil {
		return err
	}

	block, err := createAES_CipherBlock(key)
	if err != nil {
		return err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	buf := cipherText[:aes.BlockSize]

	if _, err := readRandBytesToBuf(buf); err != nil {
		return err
	}

	mode := cipher.NewCBCEncrypter(block, buf)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return os.WriteFile(outputFile, cipherText, 0644)
}

// DecryptAES_File decrypts a file using AES algorithm if it is available.
// It returns any error that may have occured.
func DecryptAES_File(inputFile, outputFile string, key []byte) error {
	cipherText, err := readInputFile(inputFile)
	if err != nil {
		return err
	}

	block, err := createAES_CipherBlock(key)
	if err != nil {
		return err
	}

	buf := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, buf)
	mode.CryptBlocks(cipherText, cipherText)

	return os.WriteFile(outputFile, cipherText, 0644)
}
