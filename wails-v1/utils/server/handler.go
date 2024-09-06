package server

/*
This module handles all elements of the server, including listening for incoming connections and handling them, decoding messages and data from clients, and printing it for the user.

Functions included in this file:
- Listen (Listen: port) --> nil
- Handle connection (handleConnection: conn) --> nil
- Handle message (handleMessage: message) --> nil
- Handle data (handleData: data) --> nil
*/

import (
	"io"
	"log"
	"myproject/utils/encryption"
	"myproject/utils/users"
	"net"
	"strings"
)

// Handle connection (handleConnection: conn) --> nil
// This function handles a connection from a client. It is called by the Listen function.
func handleConnection(conn net.Conn, self_privateKey string) {
	log.Println("Connection established")
	defer conn.Close()

	// read data from the connection
	data, err := io.ReadAll(conn)
	if err != nil {
		log.Println("Error reading data:", err.Error())
		return
	}
	log.Println("Data received:", string(data))

	// split the data into the username and encrypted data they are separated by a comma
	splitData := strings.Split(string(data), ",")
	username := splitData[0]
	encryptedData := splitData[1]

	// decrypt the data
	decryptedData, err := encryption.Decrypt([]byte(encryptedData), []byte(self_privateKey))
	if err != nil {
		log.Println("Error decrypting data:", err.Error())
		return
	}

	address := conn.RemoteAddr().String()

	// if the decrypted data begins with `message:`, handle it as a message
	if strings.HasPrefix(string(decryptedData), "message:") {
		handleMessage(username, decryptedData, address)
	} else if strings.HasPrefix(string(decryptedData), "data:") {
		handleData(username, decryptedData, address)
	} else {
		log.Println("Data not recognised")
	}
}

// Handle data (handleData: data) --> nil
// This function handles the data received from the client. It is called by the handleConnection function.
func handleMessage(username string, decryptedData []byte, address string) {
	log.Println("Message received from ", username, "@"+address+" : ", string(decryptedData)[8:])
}

// Handle data (handleData: data) --> nil
// This function handles the data received from the client. It is called by the handleConnection function.
func handleData(username string, decryptedData []byte, address string) {
	// get the datatype of the data, first strip teh `data:` from the beginning of the data then check the first word of the data to match user_info, peers, users.
	data := strings.Split(string(decryptedData)[5:], " ")
	if data[0] == "user_info" {
		// send the user info to the users module, first check if we already have the user info, then check if the user received is the same as the user we have info for, then append the user info to the user info list
		log.Println("User info received from", username, ":", strings.Join(data[1:], " "))

		remote_username := data[1]
		remote_publicKey := data[2]

		// check if the user is validated by the peer
		if users.ValidateUser(remote_username, remote_publicKey) {
			log.Println("User", remote_username, "validated by peer")
			return
		} else {
			log.Println("User", remote_username, "not validated by peer")
			// get the user info from the user providing the user info
			// peer_username, peer_publicKey := client.GetUserInfo(address)
			peer_username, peer_publicKey := "peer_username", "peer_publicKey"
			// validate the user info
			if users.ValidateUser(peer_username, peer_publicKey) {
				// get the authority of the user providing the user info
				peer_authority := users.GetAuthority(peer_username)
				// get the authority of the user info provided by the peer
				remote_authority := users.GetAuthority(remote_username)
				// if the authority of the peer is greater than or equal to the authority of the user providing the user info, update the user info
				if peer_authority >= remote_authority {
					users.UpdateUser(remote_username, remote_publicKey, "")
					log.Println("User info updated for", remote_username)
				} else {
					log.Println("User info not updated for" + remote_username + " due to insufficient authority from peer " + peer_username)
				}
			} else {
				log.Println("Peer " + address + " does not have valid user info for " + remote_username)
			}
		}
	}
	if data[0] == "peers" {
		// split the payload after removing the `peers` keyword by each , and send the peers to the client
		log.Println("Peers received from", username, ":", strings.Join(data[1:], ", "))

		for _, peer := range data[1:] {
			// send the peer to the peers module
			// client.AddPeer(peer)
			log.Println("Peer added:", peer)
		}
	}
	if data[0] == "users" {
		log.Println("Users received from", username, ":", strings.Join(data[1:], " "))

		// the data is in the fomat `users user1,publickey1 user2,publickey2 ...`
		for _, user := range data[1:] {
			// split the user into username and public key
			user_info := strings.Split(user, ",")
			remote_username := user_info[0]
			remote_publicKey := user_info[1]

			// validate the user
			if users.ValidateUser(remote_username, remote_publicKey) {
				log.Println("User", remote_username, "validated by peer")
				continue
			} else {
				log.Println("User", remote_username, "not validated by peer")

				local_publicKey, _ := users.GetUser(remote_username)
				if local_publicKey != "" {
					// send the user info to the users module to save the user info
					users.CreateUser(remote_username, remote_publicKey, "")
				} else {

					// get the user info from the user providing the user info
					// peer_username, peer_publicKey := client.GetUserInfo(address)
					peer_username, peer_publicKey := "peer_username", "peer_publicKey"
					// validate the user info
					if users.ValidateUser(peer_username, peer_publicKey) {
						// get the authority of the user providing the user info
						peer_authority := users.GetAuthority(peer_username)
						// get the authority of the user info provided by the peer
						remote_authority := users.GetAuthority(remote_username)
						// if the authority of the peer is greater than or equal to the authority of the user providing the user info, update the user info
						if peer_authority >= remote_authority {
							users.UpdateUser(remote_username, remote_publicKey, "")
							log.Println("User info updated for", remote_username)
						} else {
							log.Println("User info not updated for" + remote_username + " due to insufficient authority from peer " + peer_username)
						}
					} else {
						log.Println("Peer " + address + " does not have valid user info for " + remote_username)
					}
				}
			}
		}
	}
}
