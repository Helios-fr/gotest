package users

/*
This modile handles all elements of user data, including the creation and replication of users from peers, the storage and retrival of user data, and the management of user data.

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
	"encoding/csv"
	"io"
	"log"
	"os"
)

// RemoveUser --> bool
// This function removes the user from the user data csv file.
func RemoveUser(username string) bool {
	// Open the user data csv file for reading
	file, err := os.Open("user_data.csv")
	if err != nil {
		log.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	// Create a temporary file to write the updated data
	tempFile, err := os.CreateTemp("", "user_data_temp.csv")
	if err != nil {
		log.Println("Error creating temporary file:", err)
		return false
	}
	defer tempFile.Close()

	// Read the user data csv file
	reader := csv.NewReader(file)
	writer := csv.NewWriter(tempFile)

	// Track if the user was found and removed
	userRemoved := false

	// Find the user in the csv file and write to the temporary file
	for {
		// Read the next line of the csv file
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			log.Println("Error reading file:", err)
			return false
		}

		// Check if the username matches the current record
		if record[0] == username {
			// Skip writing this record to effectively remove the user
			userRemoved = true
			continue
		}

		// Write the record to the temporary file
		if err := writer.Write(record); err != nil {
			log.Println("Error writing to temporary file:", err)
			return false
		}
	}

	// Flush the writer
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Println("Error flushing writer:", err)
		return false
	}

	// Close the files
	file.Close()
	tempFile.Close()

	// Replace the original file with the temporary file
	if err := os.Rename(tempFile.Name(), "user_data.csv"); err != nil {
		log.Println("Error replacing original file:", err)
		return false
	}

	return userRemoved
}
