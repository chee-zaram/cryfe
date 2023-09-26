package encryptor

import (
	"os"
	"testing"
)

func Test_readInputFile(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Write some data to the temporary file
	data := []byte("test data")
	if _, err := tmpfile.Write(data); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Call the function under test
	result, err := readInputFile(tmpfile.Name())
	// Check if there was an error
	if err != nil {
		t.Errorf("readInputFile returned an error: %v", err)
	}

	// Check if the returned data matches the expected data
	if string(result) != string(data) {
		t.Errorf("readInputFile returned incorrect data, got: %s, want: %s", string(result), string(data))
	}
}

func Test_readInputFile_NonExistentFile(t *testing.T) {
	// Call the function under test with a non-existent file
	_, err := readInputFile("nonexistent.txt")
	if err == nil {
		t.Error("readInputFile did not return an error for a non-existent file")
	}
}

func Test_createAES_CipherBlock(t *testing.T) {
	_, err := createAES_CipherBlock(nil)
	if err == nil {
		t.Errorf("createAES_CipherBlock failed to return an error")
	}
}
