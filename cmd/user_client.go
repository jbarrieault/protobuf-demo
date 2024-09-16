package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"

	"github.com/jbarrieault/protobuf-demo/pkg/user"
	"google.golang.org/protobuf/proto"
)

const socketPath = "/tmp/protobuf-demo-socket.sock"

func main() {
	u := &user.User{
		Id:        71872,
		FirstName: "Jack",
		LastName:  "Barrieault",
	}

	data, err := proto.Marshal(u)
	if err != nil {
		log.Fatalf("Failed to serialize user: %v", err)
	}
	fmt.Printf("Serialized user: %x\n", data)

	encoded := base64.StdEncoding.EncodeToString(data)

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to connect to socket: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(encoded))
	if err != nil {
		log.Fatalf("failed to write to socket: %v", err)
	}

	fmt.Println("User written to socket!")
}
