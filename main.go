package main

import (
	"fmt"
	"log"

	"github.com/jbarrieault/protobuf-demo/pkg/user"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("hello, let's learn some protobuf")

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

	u2 := &user.User{}
	err = proto.Unmarshal(data, u2)
	if err != nil {
		log.Fatalf("Failed to deserialize user: %v", err)
	}

	fmt.Printf("u:  %v\n", u)
	fmt.Printf("u2: %v\n", u2)
}
