package main

import (
	"fmt"
	"log"
	"protobuf/pb"

	"google.golang.org/protobuf/proto"
)

func main() {

	// Create new expression
	person := &pb.Person{
		Name:    "John Doe",
		Age:     26,
		Address: "New York City",
	}

	// Protobuf Marshal person to out
	out, err := proto.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
	// Output: [10 8 74 111 104 110 32 68 111 101 16 26 26 13 78 101 119 32 89 111 114 107 32 67 105 116 121]

	// Protobuf Unmarshal out to person (in)
	in := &pb.Person{}
	if err := proto.Unmarshal(out, in); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(in)
	// Output: name:"John Doe" age:26 address:"New York City"

}
