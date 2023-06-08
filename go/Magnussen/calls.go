package Magnussen

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) SendCommand(command string, args ...string) (string, error) {
	if c.conn == nil {
		return "", errors.New("Not connected to Appledore")
	}

	redisCommand := c.buildCommand(command, args)
	_, err := fmt.Fprint(c.conn, redisCommand)
	if err != nil {
		return "", err
	}
	responseData := ""
	reader := bufio.NewReader(c.conn)
	// TODO: Add support for data > 1024 bytes
	buf := make([]byte, 1024)
	_, err = reader.Read(buf)
	if err != nil {
		panic(err)

	}
	responseData += c.encoder(strings.Split(string(buf), "\r\n")[len(strings.Split(string(buf), "\r\n"))-2])

	return responseData, nil
}

func (c *Client) buildCommand(command string, args []string) string {
	totalArgs := 1 + len(args)

	redisCommand := fmt.Sprintf("*%d\r\n$%d\r\n%s\r\n", totalArgs, len(command), command)

	for _, arg := range args {
		redisCommand += fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
	}

	return redisCommand
}

func (c *Client) encoder(data string) string {
	noEscapeString := strings.ReplaceAll(data, "\r\n", "")
	noEscapeString = strings.ReplaceAll(noEscapeString, "$\\d+", "")

	if strings.HasPrefix(noEscapeString, "+") {
		noEscapeString = noEscapeString[1:]
	}

	if strings.HasPrefix(noEscapeString, ":") {
		noEscapeString = noEscapeString[1:]
	}

	return noEscapeString
}
