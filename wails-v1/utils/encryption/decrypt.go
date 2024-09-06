package encryption

/*
This  module handles all elements of the encryption and decryption of the data, enabling the main function to deal with data as if it is not encrypted.

Functions included in this file:
- Generate key pair (GenerateKeyPair) --> publicKey, privateKey
- Encrypt data (Encrypt: data, publicKey) --> encryptedData
- Decrypt data (Decrypt: encryptedData, privateKey) --> data
*/

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// Decrypt --> encryptedData, privateKey --> data
// This function decrypts the encrypted data using the private key.
func Decrypt(encryptedData, privateKey []byte) (data []byte, err error) {
	// Decode private key
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, err
	}
	decodedPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Decrypt data
	data, err = rsa.DecryptPKCS1v15(rand.Reader, decodedPrivateKey, encryptedData)
	if err != nil {
		return nil, err
	}

	return data, nil
}
