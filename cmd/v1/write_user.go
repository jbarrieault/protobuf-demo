package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"

	"github.com/jbarrieault/protobuf-demo/pkg/user"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const socketPath = "/tmp/protobuf-demo-socket.sock"

func main() {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to connect to socket: %v\nIs user_server.rb running?", err)
	}
	defer conn.Close()

	u := &user.User{
		WriterVersion: 1,
		Id:            71872,
		FirstName:     "Jack",
		LastName:      "Black",
		MiddleName:    "Jacob",
		// The v2 schema changed emails type,
		// which means adding this will break any reads that use the v2 schema
		// Email:         "jack@example.com",
	}

	_, err = conn.Write(base64ProtoMessageBytes(u))
	if err != nil {
		log.Fatalf("failed to write to socket: %v", err)
	}

	fmt.Println("User message written to socket!")
}

func base64ProtoMessageBytes(m protoreflect.ProtoMessage) []byte {
	data, err := proto.Marshal(m)
	if err != nil {
		log.Fatalf("Failed to serialize message: %v", err)
	}
	fmt.Printf("Serialized user: %x\n", data)

	return []byte(base64.StdEncoding.EncodeToString(data))
}
