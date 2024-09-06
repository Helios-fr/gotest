## Server Module
This module handles all elements of the server, including listening for incoming connections and handling them, decoding messages and data from clients, and printing it for the user.

Functions included in this file:
- Listen (Listen: port) --> nil
- Handle connection (handleConnection: conn) --> nil
- Handle message (handleMessage: message) --> nil
- Handle data (handleData: data) --> nil

### Listen --> nil
This function listens for incoming connections on a specified port. It is called by the main function in the server module.

### HandleConnection --> nil
This function handles a connection from a client. It is called by the Listen function.

### HandleMessage --> nil
This function handles a message for the client, using the users module to verify the user, then the decryption module to decrypt the message.

### HandleData --> nil
This function handles data from peers, using the users module to verify the user, then the decryption module to decrypt the data and store it in the database.