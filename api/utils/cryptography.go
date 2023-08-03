package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"
)

func LoadPrivateKey(fileName string) (*rsa.PrivateKey, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if file != nil {
			if err := file.Close(); err != nil {
				fmt.Println("Error closing files:", err)
			}
		}
	}()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(fileData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	actualPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("AAA")
	}

	return actualPrivateKey, nil
}

func LoadPublicKeyFromFile(fileName string) (*rsa.PublicKey, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if file != nil {
			if err := file.Close(); err != nil {
				fmt.Println("Error closing files:", err)
			}
		}
	}()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(fileData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key")
	}

	return publicKey, nil
}

func RsaEncrypt(publicKey *rsa.PublicKey, plainText []byte) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, plainText, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func RsaDecrypt(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	plainText, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

func RsaDecryptBase64(privateKey *rsa.PrivateKey, base64Ciphertext string) ([]byte, error) {

	ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil {
		return nil, err
	}

	plainText, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func EncryptECB(key, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Pad the plaintext to be a multiple of AES block size
	blockSize := block.BlockSize()
	padding := blockSize - len(plaintext)%blockSize
	paddedPlaintext := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)

	ciphertext := make([]byte, len(paddedPlaintext))
	block.Encrypt(ciphertext, paddedPlaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil

}

func DecryptECB(key []byte, ciphertext string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	// ECB mode does not need an IV, so we use a nil IV.
	plaintext := make([]byte, len(decodedCiphertext))
	block.Decrypt(plaintext, decodedCiphertext)

	return plaintext, nil
}
func test() {
	privateKey, err := LoadPrivateKey("secret_key/private_key.pem")
	if err != nil {
		fmt.Println("Error loading private key:", err)
		return
	}
	publicKey, err := LoadPublicKeyFromFile("secret_key/public_key.pem")
	if err != nil {
		fmt.Println("Error loading public key:", err)
		return
	}

	// 加密和解密
	plainText := []byte("Hello, RSA OAEP SHA1")
	ciphertext, err := RsaEncrypt(publicKey, plainText)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}

	decryptedText, err := RsaDecrypt(privateKey, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}

	fmt.Println("Original text:", string(plainText))
	fmt.Println("Decrypted text:", string(decryptedText))
	key := []byte("0123456789abcdef")    // 16-byte AES key
	plaintext := []byte("Hello, World!") // Plain text to encrypt

	encrypted, err := EncryptECB(key, plaintext)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	fmt.Println("Encrypted:", encrypted)

	decrypted, err := DecryptECB(key, encrypted)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	fmt.Println("Decrypted:", string(decrypted))

}
