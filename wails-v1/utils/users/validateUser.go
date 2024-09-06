package users

import "strings"

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

// ValidateUser --> bool
// This function validates the user by checking the public key against the stored public key for the given username.
func ValidateUser(username string, publicKey string) bool {
	pub, _ := GetUser(username)
	publicKey = strings.ReplaceAll(publicKey, "\n", "")

	return pub == publicKey
}
