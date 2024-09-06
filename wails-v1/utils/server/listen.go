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
	"log"
	"net"
)

// Listen --> server
// This function listens for incoming connections on the specified port.
func Listen(port string, self_privateKey string) {
	// create a listener
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	log.Println("Listening on " + port)

	// listen for incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			return
		}
		go handleConnection(conn, self_privateKey)
	}
}
