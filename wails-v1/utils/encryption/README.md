## Encryption Module
This  module handles all elements of the encryption and decryption of the data, enabling the main function to deal with data as if it is not encrypted.

Functions included in this file:
- Generate key pair (GenerateKeyPair) --> publicKey, privateKey
- Encrypt data (Encrypt: data*, publicKey*) --> encryptedData
- Decrypt data (Decrypt: encryptedData*, privateKey*) --> data

### GenerateKeyPair --> publicKey, privateKey
This function generates a public and private key pair for the encryption and decryption of data.

### Encrypt --> data, publicKey --> encryptedData
This function encrypts the data using the public key.

### Decrypt --> encryptedData, privateKey --> data
This function decrypts the encrypted data using the private key.