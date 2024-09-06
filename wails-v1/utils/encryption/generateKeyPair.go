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

// GenerateKeyPair --> publicKey, privateKey
// This function generates a public and private key pair for the encryption and decryption of data.
func GenerateKeyPair() (publicKey, privateKey []byte, err error) {
	// Generate key pair
	keyPair, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Encode public key
	publicKey = x509.MarshalPKCS1PublicKey(&keyPair.PublicKey)
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKey,
	}
	publicKey = pem.EncodeToMemory(publicKeyBlock)

	// Encode private key
	privateKey = x509.MarshalPKCS1PrivateKey(keyPair)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKey,
	}
	privateKey = pem.EncodeToMemory(privateKeyBlock)

	return publicKey, privateKey, nil
}
