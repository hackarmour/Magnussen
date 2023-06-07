package Magnussen

import (
	"fmt"
	"net"
)

func Connect(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	client := NewClient(conn)
	return client, nil
}
