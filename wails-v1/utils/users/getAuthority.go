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
- Get Authority (GetAuthority: username*) --> bool
*/

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// GetAuthority --> int
// This function returns the line index of a given username in the user data csv file.
func GetAuthority(username string) int {
	// Open the user data csv file
	file, err := os.Open("user_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the user data csv file
	reader := csv.NewReader(file)
	reader.Comment = '#'

	// Search for the user in the csv file
	for i := 0; ; i++ {
		// Read the next line of the csv file
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			log.Fatal(err)
		}

		// Check if the username matches the current record
		if record[0] == username {
			// Return the line index of the user
			return i
		}
	}

	// Return -1 if the user is not found
	return -1
}
