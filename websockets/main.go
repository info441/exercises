package main

import (
	"fmt"
	"log"
	"net/http"
  "sync"
	"github.com/gorilla/websocket"
)

// A simple store to store all the connections
type socketStore struct {
	Connections []*websocket.Conn
  lock sync.Mutex
}

// Control messages for websocket
const (
	// TextMessage denotes a text data message. The text message payload is
	// interpreted as UTF-8 encoded text data.
	TextMessage = 1

	// BinaryMessage denotes a binary data message.
	BinaryMessage = 2

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)

// Thread-safe method for inserting a connection
func (s *socketStore) InsertConnection(conn *websocket.Conn) int {
  s.lock.Lock()
  connId := len(s.Connections)
  // insert socket connection
	s.Connections = append(s.Connections, conn)
  s.lock.Unlock()
  return connId
}

// Thread-safe method for inserting a connection
func (s *socketStore) RemoveConnection(connId int) {
  s.lock.Lock()
  // insert socket connection
	s.Connections = append(s.Connections[:connId], s.Connections[connId+1:]...)
  s.lock.Unlock()
}

// Simple method for writing a message to all live connections.
// In your homework, you will be writing a message to a subset of connections
// (if the message is intended for a private channel), or to all of them (if the message
// is posted on a public channel
func (s *socketStore) WriteToAllConnections(messageType int, data []byte) error {
  var writeError error;

  for _, conn := range s.Connections {
    writeError = conn.WriteMessage(messageType, data)
    if writeError != nil {
      return writeError
    }
  }

  return nil
}


// This is a struct to read our message into
type msg struct {
	Message string `json:"message"`
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
    // This function's purpose is to reject websocket upgrade requests if the
    // origin of the websockete handshake request is coming from unknown domains.
    // This prevents some random domain from opening up a socket with your server.
    // TODO: make sure you modify this for your HW to check if r.Origin is your host
		return true
	},
}


func (sockets *socketStore) webSocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
	// handle the websocket handshake
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to open websocket connection", 401)
		return
	}

  // Insert our connection onto our datastructure for ongoing usage
  connId := sockets.InsertConnection(conn)

  // Invoke a goroutine for handling control messages from this connection
  go (func(conn *websocket.Conn, connId int) {
    defer conn.Close()
    defer sockets.RemoveConnection(connId)

    for {
      messageType, p, err := conn.ReadMessage()

      if messageType == TextMessage || messageType == BinaryMessage {
        fmt.Printf("Client says %v\n", p)
        fmt.Printf("Writing %s to all sockets\n", string(p))
        sockets.WriteToAllConnections(TextMessage, append([]byte("Hello from server: "), p...))
      } else if messageType == CloseMessage {
        fmt.Println("Close message received.")
        break
      } else if err != nil {
        fmt.Println("Error reading message.")
        break
      }
      // ignore ping and pong messages
    }

  })(conn, connId)
}


func main() {
	mux := http.NewServeMux()

	ctx := socketStore{
		Connections: []*websocket.Conn{},
	}

	mux.HandleFunc("/ws", ctx.webSocketConnectionHandler)
	log.Fatal(http.ListenAndServe(":4001", mux))
}
