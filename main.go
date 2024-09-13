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
		FirstName: "Jack",
		LastName:  "Barrieault",
	}

	data, err := proto.Marshal(u)
	if err != nil {
		log.Fatalf("Failed to serialize user: %v", err)
	}
	fmt.Printf("Serialized user: %x\n", data)

}
