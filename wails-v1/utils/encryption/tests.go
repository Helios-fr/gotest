package encryption

/*
This  module handles all elements of the encryption and decryption of the data, enabling the main function to deal with data as if it is not encrypted.

Functions included in this file:
- Generate key pair (GenerateKeyPair) --> publicKey, privateKey
- Encrypt data (Encrypt: data, publicKey) --> encryptedData
- Decrypt data (Decrypt: encryptedData, privateKey) --> data
*/

// Test_GenerateKeyPair --> publicKey, privateKey
// This function tests the generation of a public and private key pair for the encryption and decryption of data.
func Test_GenerateKeyPair() (success bool) {
	_, _, err := GenerateKeyPair()
	return err == nil
}

// Test_Encrypt --> data, publicKey --> encryptedData
// This function tests the encryption of the data using the public key.
func Test_Encrypt() (success bool) {
	if !Test_GenerateKeyPair() {
		return false
	}

	publicKey, _, _ := GenerateKeyPair()

	_, err := Encrypt([]byte("test"), publicKey)
	return err == nil
}

// Test_Decrypt --> encryptedData, privateKey --> data
// This function tests the decryption of the encrypted data using the private key.
func Test_Decrypt() (success bool) {
	if !Test_GenerateKeyPair() {
		return false
	}

	publicKey, privateKey, _ := GenerateKeyPair()

	encryptedData, _ := Encrypt([]byte("test"), publicKey)

	_, err := Decrypt(encryptedData, privateKey)
	return err == nil
}

// Test_All --> bool
// This function tests all functions in the encryption module.
func Test_All() (success bool) {
	return Test_GenerateKeyPair() && Test_Encrypt() && Test_Decrypt()
}
