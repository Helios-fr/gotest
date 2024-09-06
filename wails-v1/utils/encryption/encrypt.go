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

// Encrypt --> data, publicKey --> encryptedData
// This function encrypts the data using the public key.
func Encrypt(data, publicKey []byte) (encryptedData []byte, err error) {
	// Decode public key
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, err
	}
	decodedPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Encrypt data
	encryptedData, err = rsa.EncryptPKCS1v15(rand.Reader, decodedPublicKey, data)
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}
