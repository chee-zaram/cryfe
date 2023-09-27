package encryptor

import (
	"crypto/aes"
	"os"
	"testing"
)

func Test_readInputFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	data := []byte("test data")
	if _, err := tmpfile.Write(data); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	result, err := readInputFile(tmpfile.Name())
	if err != nil {
		t.Errorf("readInputFile returned an error: %v", err)
	}

	if string(result) != string(data) {
		t.Errorf("readInputFile returned incorrect data, got: %s, want: %s", string(result), string(data))
	}
}

func Test_readInputFile_NonExistentFile(t *testing.T) {
	_, err := readInputFile("nonexistent.txt")
	if err == nil {
		t.Error("readInputFile did not return an error for a non-existent file")
	}
}

func Test_createAES_CipherBlock(t *testing.T) {
	key := make([]byte, 16) // AES-128 key

	block, err := createAES_CipherBlock(key)
	if err != nil {
		t.Fatalf("Error creating AES cipher block: %v", err)
	}
	if block.BlockSize() != aes.BlockSize {
		t.Errorf("Expected block size %d, but got %d", aes.BlockSize, block.BlockSize())
	}
}

func Test_createAES_CipherBlock_NilKey(t *testing.T) {
	_, err := createAES_CipherBlock(nil)
	if err == nil {
		t.Errorf("Expected an error for a nil key, but got nil")
	}
}

func Test_readRandBytesToBuf(t *testing.T) {
	buf := make([]byte, 16)

	n, err := readRandBytesToBuf(buf)
	if err != nil {
		t.Fatalf("Error reading random bytes: %v", err)
	}
	if n != len(buf) {
		t.Errorf("Expected to read %d bytes, but got %d", len(buf), n)
	}
}

func Test_readRandBytesToBuf_NilBuffer(t *testing.T) {
	n, _ := readRandBytesToBuf(nil)
	if n != 0 {
		t.Errorf("Expected number of bytes read to be 0, got %d", n)
	}
}

/* func TestEncryptFileAndDecryptFile(t *testing.T) {
	key := make([]byte, 16) // AES-128 key

	tmpFile := createTempFile(t)
	defer os.Remove(tmpFile.Name())

	plainText := []byte("Testcontent")
	if err := os.WriteFile(tmpFile.Name(), plainText, 0644); err != nil {
		t.Fatal(err)
	}

	encryptedFile := tmpFile.Name() + ".enc"
	err := EncryptFile(tmpFile.Name(), encryptedFile, key)
	if err != nil {
		t.Fatalf("Error encrypting file: %v", err)
	}

	decryptedFile := tmpFile.Name() + ".dec"
	err = DecryptAES_File(encryptedFile, decryptedFile, key)
	if err != nil {
		t.Fatalf("Error decrypting file: %v", err)
	}

	decryptedContent, err := os.ReadFile(decryptedFile)
	if err != nil {
		t.Fatalf("Error reading decrypted file: %v", err)
	}

	if string(plainText) != string(decryptedContent) {
		t.Errorf("Decrypted content does not match original content")
	}
} */

// createTempFile creates a temporary test file and returns a file handle.
func createTempFile(t *testing.T) *os.File {
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Error creating temporary test file: %v", err)
	}
	return tmpFile
}
