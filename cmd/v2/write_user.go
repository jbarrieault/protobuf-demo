package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"

	user "github.com/jbarrieault/protobuf-demo/pkg/user-v2"

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
		WriterVersion: 2,
		Id:            71872,
		FirstName:     "Jack",

		// clients using v1 schema correctly map this to last_name field :)
		Surname: "Black",

		// clients using v1 schema _almost_ seamlessly handles the type change,
		// it's coming out the other side with leading characters:
		// "\n\u0016jack.black@example.com"
		Email: &user.Email{Address: "jack@example.com"},

		// clients using v1 schema don't know about this field, it does not appear
		Age: 55,
	}

	_, err = conn.Write(base64ProtoMessageBytes(u))
	if err != nil {
		log.Fatalf("failed to write to socket: %v", err)
	}

	fmt.Println("V2 User written to socket!")
}

func base64ProtoMessageBytes(m protoreflect.ProtoMessage) []byte {
	data, err := proto.Marshal(m)
	if err != nil {
		log.Fatalf("Failed to serialize message: %v", err)
	}
	fmt.Printf("Serialized user: %x\n", data)

	return []byte(base64.StdEncoding.EncodeToString(data))
}
