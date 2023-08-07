package main

import (
	"fmt"
	"log"

	"github.com/hackarmour/Magnussen/Magnussen"
)

func main() {
	client, err := Magnussen.Connect("localhost:6379")
	if err != nil {
		log.Fatal(err.Error())
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
	lsetResponse, err := client.SendCommand("lset", "myarr", "0", "myvalue")
	if err != nil {
		log.Fatal("Failed to send LSET command:", err)
	}

	lpushResponse, err := client.SendCommand("lpush", "myarr", "myvalue2")
	if err != nil {
		log.Fatal("Failed to send LPUSH command:", err)
	}
	rpushResponse, err := client.SendCommand("rpush", "myarr", "myvalue3")
	if err != nil {
		log.Fatal("Failed to send RPUSH command:", err)
	}
	llenResponse, err := client.SendCommand("llen", "myarr")
	if err != nil {
		log.Fatal("Failed to send LLEN command:", err)
	}
	lpopResponse, err := client.SendCommand("lpop", "myarr")
	if err != nil {
		log.Fatal("Failed to send LPOP command:", err)
	}
	lrangeResponse, err := client.SendCommand("lrange", "myarr", "0", "3")
	if err != nil {
		log.Fatal("Failed to send LRANGE command:", err)
	}
	fmt.Println("SET:", setKeyResponse)
	fmt.Println("PING:", pingResponse)
	fmt.Println("ECHO:", echoResponse)
	fmt.Println("GET:", getKeyResponse)
	fmt.Println("DEL:", deleteKeyResponse)
	fmt.Println("LSET:", lsetResponse)
	fmt.Println("LPUSH:", lpushResponse)
	fmt.Println("RPUSH:", rpushResponse)
	fmt.Println("LLEN:", llenResponse)
	fmt.Println("LPOP:", lpopResponse)
	fmt.Println("LRANGE:", lrangeResponse)

}
