package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("Hello world")
	Alexander := &Person{
		Name:    "Berezovskiy",
		Age:     31,
		Friends: &Friends{RealFriends: 5},
	}
	data, err := proto.Marshal(Alexander)
	if err != nil {
		log.Fatal("error")
	}
	fmt.Println(data)
	newAlexander := &Person{}
	err = proto.Unmarshal(data, newAlexander)
	if err != nil {
		log.Fatal("failed unmarshalling", err)
	}
	fmt.Println(newAlexander.GetName())
	fmt.Println(newAlexander.GetAge())
	fmt.Println(newAlexander.Friends.GetRealFriends())

}
