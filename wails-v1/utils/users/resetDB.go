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
	"log"
	"os"
)

// ResetDB --> bool
// This function resets the user data csv file, removing all user data.
func ResetDB() bool {
	// Create a new user data csv file to overwrite the existing file
	file, err := os.Create("user_data.csv")
	if err != nil {
		log.Println("Error creating file:", err)
		return false
	}
	defer file.Close()

	return true
}
