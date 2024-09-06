package users

/*
This module handles all elements of user data, including the creation and replication of users from peers, the storage and retrieval of user data, and the management of user data.

Functions included in this file:
- Create user (CreateUser: username*, publicKey*, privateKey) --> bool
- Get user (GetUser: username*) --> publicKey, privateKey
- Update user (UpdateUser: username*, publicKey*, privateKey) --> bool
- Validate user (ValidateUser: username*, publicKey) --> bool
- Remove user (RemoveUser: username*) --> bool
- Reset DB (resetDB: ) --> bool
- Get Authority (GetAuthority: username*) --> int
*/

import (
	"log"
	"os"
	"strings"
)

// UpdateUser --> bool
// This function updates the public key and private key for a given username.
func UpdateUser(username, publicKey, privateKey string) bool {
	// Read the entire user data CSV file
	data, err := os.ReadFile("user_data.csv")
	if err != nil {
		log.Fatal(err)
		return false
	}

	// remove the new line character from the public and private keys
	// replace the new line character with an empty string
	publicKey = strings.ReplaceAll(publicKey, "\n", "")
	privateKey = strings.ReplaceAll(privateKey, "\n", "")

	// Split the file into lines
	lines := strings.Split(string(data), "\n")

	// Find and update the user entry
	updated := false
	for i, line := range lines {
		// Split the line into fields
		fields := strings.Split(line, ",")
		if len(fields) < 3 {
			continue // Skip malformed lines
		}

		// Update the record if the username matches
		if fields[0] == username {
			fields[1] = publicKey
			fields[2] = privateKey
			lines[i] = strings.Join(fields, ",")
			updated = true
			break
		}
	}

	if !updated {
		log.Println("User not found")
		return false
	}

	// Join the lines back into a single string
	output := strings.Join(lines, "\n")

	// Write the updated data back to the file
	err = os.WriteFile("user_data.csv", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
