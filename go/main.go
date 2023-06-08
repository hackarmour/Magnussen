package main

import (
	"fmt"
	"log"

	"github.com/hackarmour/Magnussen/Magnussen"
)

func main() {
	client, err := Magnussen.Connect("localhost:6379")
	if err != nil {
		log.Fatal("Failed to connect to Appledore:", err)
	}

	fmt.Println("Connected to Appledore")

	setKeyResponse, err := client.SendCommand("set", "mykey", "myvalue")
	if err != nil {
		log.Fatal("Failed to send SET command:", err)
	}

	pingResponse, err := client.SendCommand("ping")
	if err != nil {
		log.Fatal("Failed to send PING command:", err)
	}

	echoResponse, err := client.SendCommand("echo", "hello world")
	if err != nil {
		log.Fatal("Failed to send ECHO command:", err)
	}

	getKeyResponse, err := client.SendCommand("get", "mykey")
	if err != nil {
		log.Fatal("Failed to send GET command:", err)
	}

	deleteKeyResponse, err := client.SendCommand("del", "mykey")
	if err != nil {
		log.Fatal("Failed to send DEL command:", err)
	}

	fmt.Println("SET:", setKeyResponse)
	fmt.Println("PING:", pingResponse)
	fmt.Println("ECHO:", echoResponse)
	fmt.Println("GET:", getKeyResponse)
	fmt.Println("DEL:", deleteKeyResponse)

}
