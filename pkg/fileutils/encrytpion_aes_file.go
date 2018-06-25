package fileutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

var keyLen = 32
var saltLen = 32

// AESEncryptFile takes a file and passphrase and encrypts the file
func AESEncryptFile(infilepath string, passphrase string, outfilepath string) {
	// Get random salt
	salt := make([]byte, saltLen)
	if _, err := rand.Reader.Read(salt); err != nil {
		panic("random reader failed")
	}
	// Derive key
	key := pbkdf2.Key([]byte(passphrase), salt, 4096, keyLen, sha1.New)

	// Byte array of the string
	plaintext, err := ioutil.ReadFile(infilepath)
	if err != nil {
		panic(err.Error())
	}

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// Empty array of 16 + plaintext length
	ciphertext := make([]byte, aes.BlockSize+len(salt)+len(plaintext))

	// The IV needs to be unique, but not secure. Therefore it's common to include it at the beginning of the ciphertext.
	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]
	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// The salt can be stored in the file as well (this isnt super secure)
	filePos := aes.BlockSize + len(salt)
	slt := ciphertext[aes.BlockSize:filePos]
	if _, err := io.ReadFull(bytes.NewReader(salt), slt); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[filePos:], plaintext)

	// create a new file for saving the encrypted data.
	f, err := os.Create(outfilepath)
	if err != nil {
		panic(err.Error())
	}

	//Copy it in
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}
}

// AESDecryptFile takes a file and passphrase and decrypts the file
func AESDecryptFile(infilepath string, passphrase string, outfilepath string) {
	// Byte array of the string
	ciphertext, err := ioutil.ReadFile(infilepath)
	if err != nil {
		panic(err.Error())
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		panic("File is too short")
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	//Get the salt
	filePos := aes.BlockSize + saltLen
	salt := ciphertext[aes.BlockSize:filePos]

	// Derive key
	key := pbkdf2.Key([]byte(passphrase), salt, 4096, keyLen, sha1.New)

	// Remove the IV and salt from the ciphertext
	ciphertext = ciphertext[filePos:]

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	// create a new file for saving the encrypted data.
	f, err := os.Create(outfilepath)
	if err != nil {
		panic(err.Error())
	}

	//Copy it in
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}

}
