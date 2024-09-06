package users

import (
	"log"
	"os"
)

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

// Test_ResetDB --> bool
// This function tests the ResetDB function by resetting the user data csv file and checking if the file was successfully reset.
func Test_ResetDB() bool {
	return ResetDB()
}

// Test_CreateUser --> bool
// This function tests the CreateUser function by creating a new user and checking if the user was successfully added to the user data csv file.
func Test_CreateUser() bool {
	if !Test_ResetDB() {
		return false
	}

	t1 := CreateUser("testuser", "testpublickey", "testprivate")
	ResetDB()
	return t1
}

// Test_UpdateUser --> bool
// This function tests the UpdateUser function by updating the public key for a known user and checking if the public key was successfully updated in the user data csv file.
func Test_UpdateUser() bool {
	if !(Test_CreateUser() && Test_ResetDB()) {
		return false
	}

	// Create the user
	CreateUser("testuser", "testpublickey", "testprivate")

	// Update the public key for the user
	t1 := UpdateUser("testuser", "newpublickey", "testprivate")

	// Get the public key for the user
	publicKey, _ := GetUser("testuser")

	// Check if the public key was updated
	t2 := publicKey == "newpublickey"

	ResetDB()
	return t1 && t2
}

// Test_GetUser --> bool
// This function tests the GetUser function by retrieving the public key for a known user and checking if the public key is correct.
func Test_GetUser() bool {
	// Test the CreateUser function first to ensure this test is valid
	if !(Test_CreateUser() && Test_ResetDB()) {
		return false
	}

	// Create the user
	CreateUser("testuser", "testpublickey", "testprivate")

	// Get the public key for a known user
	publicKey, _ := GetUser("testuser")

	// Check if the public key is correct
	t1 := publicKey == "testpublickey"

	// Get the public key for an unknown user
	publicKey, _ = GetUser("unknownuser")

	// Check if the public key is empty
	t2 := publicKey == ""

	ResetDB()
	return t1 && t2
}

// Test_ValidateUser --> bool
// This function tests the ValidateUser function by validating a known user with the correct public key.
func Test_ValidateUser() bool {
	if !(Test_CreateUser() && Test_ResetDB()) {
		return false
	}
	// Validate the user with the correct public key
	username := "testuser"
	publicKey := "testpublickey"

	// Create the user
	CreateUser(username, publicKey, "testprivate")

	// Check if the user is validated
	t1 := ValidateUser(username, publicKey)

	publicKey = "invalidpublickey"

	// Check if the user is not validated
	t2 := !ValidateUser(username, publicKey)

	ResetDB()
	return t1 && t2
}

// Test_RemoveUser --> bool
// This function tests the RemoveUser function by removing a known user and checking if the user was successfully removed from the user data csv file.
func Test_RemoveUser() bool {
	if !(Test_CreateUser() && Test_ResetDB() && Test_GetUser()) {
		return false
	}

	// Create the user
	CreateUser("testuser", "testpublickey", "testprivate")

	// Remove the user
	t1 := RemoveUser("testuser")

	// Check if the user was removed
	_, privateKey := GetUser("testuser")
	t2 := privateKey == ""

	ResetDB()
	return t1 && t2
}

// Test_GetAuthority --> bool
// This function tests the GetAuthority function by retrieving the authority level for a known user and checking if the authority level is correct.
func Test_GetAuthority() bool {
	// Test the CreateUser function first to ensure this test is valid
	if !(Test_CreateUser() && Test_ResetDB()) {
		return false
	}

	// Create the user
	CreateUser("testuser", "testpublickey", "testprivate")

	// Get the authority level for a known user
	authority := GetAuthority("testuser")

	// Check if the authority level is correct
	t1 := authority == 0

	// Get the authority level for an unknown user
	authority = GetAuthority("unknownuser")

	// Check if the authority level is -1
	t2 := authority == -1

	ResetDB()
	return t1 && t2
}

// Test_All --> bool
// This function tests all functions in the users module.
func Test_All() bool {
	// if the user data csv file does not exist, create it
	if _, err := os.Stat("user_data.csv"); os.IsNotExist(err) {
		file, err := os.Create("user_data.csv")
		if err != nil {
			log.Println("Error creating file:", err)
			return false
		}
		file.Close()
	}

	// copy the user data csv file to a backup file
	os.Rename("user_data.csv", "user_data.csv.bak")
	t1 := Test_ResetDB() && Test_CreateUser() && Test_GetUser() && Test_ValidateUser() && Test_RemoveUser() && Test_GetAuthority() && Test_UpdateUser()
	os.Remove("user_data.csv")
	os.Rename("user_data.csv.bak", "user_data.csv")
	return t1
}
