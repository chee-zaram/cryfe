package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/chee-zaram/cryfe/encryptor"
)

var (
	cryfeCMD   = flag.NewFlagSet("cryfe", flag.ExitOnError)
	encryptCMD = flag.NewFlagSet("encrypt", flag.ExitOnError)
	decryptCMD = flag.NewFlagSet("decrypt", flag.ExitOnError)
)

var (
	encryptionKeyFile string
	encryptionKey     []byte
	decryptionKeyFile string
	decryptionKey     []byte
	inputFile         string
	outputFile        string
)

const (
	// encrypt serves as a constant string for the encrypt command.
	encrypt string = "encrypt"
	// decrypt serves as a constant string for the decrypt command.
	decrypt string = "decrypt"
)

// cryfeUsage prints usage instructions for the cryfe binary.
// Usage is printed to STDERR.
func cryfeUsage() {
	fmt.Fprintf(os.Stderr, `
CryFE encrypts and decrypts files.

    cryfe <command> [options]

For example, to encrypt a file:

    cryfe encrypt -key <your-key> <input-file> <output-file>

The following commands are supported:

    encrypt     encrypt a file
    decrypt     decrypt an encrypted file
    version     print the version of the cryfe binary
    help        print help information

To learn more about a command, run "cryfe help <command>".
`[1:])
	fmt.Fprintf(os.Stderr, `

For more information, see https://github.com/chee-zaram/cryfe.
`[1:])
}

// encryptUsage prints usage instructions for the encrypt command.
// Usage is printed to STDERR.
func encryptUsage() {
	fmt.Fprintf(os.Stderr, `
Encrypt command encrypts a file.

The following flags are supported:

`[1:])
	encryptCMD.PrintDefaults()

	fmt.Fprintf(os.Stderr, `

For more information, see https://github.com/chee-zaram/cryfe.
`[1:])
}

// decryptUsage prints usage instructions for the decrypt command.
// Usage is printed to STDERR.
func decryptUsage() {
	fmt.Fprintf(os.Stderr, `
Decrypt command decrypts a file.

The following flags are supported:

`[1:])
	decryptCMD.PrintDefaults()

	fmt.Fprintf(os.Stderr, `

For more information, see https://github.com/chee-zaram/cryfe.
`[1:])
}

// parseCLArgs parses all arguments from the command line.
func parseCLArgs() string {
	if len(os.Args) < 2 {
		cryfeUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case encrypt:
		encryptCMD.Parse(os.Args[2:])
		if encryptionKeyFile == "" {
			fmt.Fprintln(os.Stderr, "Please, provide an encryption key with -key")
			os.Exit(1)
		}

		var err error
		encryptionKey, err = os.ReadFile(encryptionKeyFile)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		if err != nil {
			encryptionKey = []byte(encryptionKeyFile)
		}

		if len(os.Args) < 5 {
			fmt.Fprintln(os.Stderr, "Please, provide file to encrypt")
			os.Exit(1)
		}
		inputFile = os.Args[4]

		if len(os.Args) < 6 {
			fmt.Fprintln(os.Stderr, "Please, provide output file name")
			os.Exit(1)
		}
		outputFile = os.Args[5]

		if len(os.Args) > 6 {
			fmt.Fprintf(os.Stderr, "\nToo many arguments to encrypt command\n\n")
			encryptUsage()
			os.Exit(1)
		}

		return encrypt

	case decrypt:
		decryptCMD.Parse(os.Args[2:])
		if decryptionKeyFile == "" {
			fmt.Fprintln(os.Stderr, "Please, provide a decryption key with -key")
			os.Exit(1)
		}

		if len(os.Args) < 5 {
			fmt.Fprintln(os.Stderr, "Please, provide file to decrypt")
			os.Exit(1)
		}
		inputFile = os.Args[4]

		if len(os.Args) < 6 {
			fmt.Fprintln(os.Stderr, "Please, provide output file name")
			os.Exit(1)
		}
		outputFile = os.Args[5]

		if len(os.Args) > 6 {
			fmt.Fprintf(os.Stderr, "\nToo many arguments to decrypt command\n\n")
			decryptUsage()
			os.Exit(1)
		}

		return decrypt

	case "help":
		if len(os.Args) == 3 {
			switch os.Args[2] {
			case encrypt:
				encryptUsage()
			case decrypt:
				decryptUsage()
			}
			os.Exit(0)
		}

		if len(os.Args) > 3 {
			fmt.Fprintln(os.Stderr, "usage: cryfe help [command]")
			os.Exit(1)
		}

		cryfeUsage()
		os.Exit(0)

	case "version":
		// TODO: implement version

	default:
		cryfeUsage()
		os.Exit(1)
	}
	return ""
}

func init() {
	cryfeCMD.Usage = cryfeUsage
	encryptCMD.Usage = encryptUsage
	encryptCMD.StringVar(&encryptionKeyFile, "key", "", "The encryption key")
	decryptCMD.StringVar(&decryptionKeyFile, "key", "", "The decryption key")
	parseCLArgs()
}

func main() {
	switch parseCLArgs() {
	case encrypt:
		if err := runEncrypt(); err != nil {
			fmt.Fprintf(os.Stderr, "encrypt: %s\n", err)
			os.Exit(2)
		}
	case decrypt:
		if err := runDecrypt(); err != nil {
			fmt.Fprintf(os.Stderr, "decrypt: %s\n", err)
			os.Exit(2)
		}
	default:
		os.Exit(0)
	}
}

// runEncrypt calls the encryption function and returns any error that may have
// occured.
func runEncrypt() error {
	return encryptor.EncryptFile(inputFile, outputFile, []byte(encryptionKey))
}

// runDecrypt calls the decryption function and returns any error that may have
// occured.
func runDecrypt() error {
	return encryptor.DecryptAES_File(inputFile, outputFile, []byte(encryptionKey))
}
