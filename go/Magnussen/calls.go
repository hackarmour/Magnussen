package Magnussen

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

const (
	maxBufferSize = 4096 // Maximum buffer size for each read iteration
)

type Client struct {
	conn  net.Conn // TCP connection to Appledore
	mutex sync.Mutex
}

// NewClient creates a new instance of the Client with the given TCP connection.
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// SendCommand sends a Redis command with the specified arguments to the Appledore server
// and returns the response received from the server.
func (c *Client) SendCommand(command string, args ...string) (string, error) {
	if c == nil || c.conn == nil {
		return "", errors.New("not connected to Appledore")
	}

	// Build the Redis command
	redisCommand := c.buildCommand(command, args)

	// Send the command to the server
	c.mutex.Lock()
	_, err := fmt.Fprint(c.conn, redisCommand)
	c.mutex.Unlock()
	if err != nil {
		return "", err
	}

	responseData := ""
	reader := bufio.NewReader(c.conn)

	// Read the response from the server
	buf, err := readLargeData(reader)
	if err != nil {
		panic(err)
	}

	// Process the response
	responseLines := strings.Split(string(buf), "\r\n")
	lastLine := responseLines[len(responseLines)-2]

	responseData += c.encoder(lastLine)
	if responseData == "" {
		return "", errors.New("empty response data")
	}
	if responseLines[0][0] == '-' {
		errMsg := responseData[1:]
		return "", errors.New(errMsg)
	} else if responseData == "$-1" {
		return "", errors.New("key not found")
	}

	return responseData, nil
}

// buildCommand builds the Redis command string from the command and arguments.
func (c *Client) buildCommand(command string, args []string) string {
	totalArgs := 1 + len(args)

	redisCommand := fmt.Sprintf("*%d\r\n$%d\r\n%s\r\n", totalArgs, len(command), command)

	for _, arg := range args {
		redisCommand += fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
	}

	return redisCommand
}

// encoder processes the data received from the server and removes escape characters.
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

// readLargeData reads large data from the reader in chunks until the end or an error occurs.
func readLargeData(reader *bufio.Reader) ([]byte, error) {
	var buf bytes.Buffer
	buffer := make([]byte, maxBufferSize)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		buf.Write(buffer[:n])

		if n < maxBufferSize {
			break
		}
	}

	return buf.Bytes(), nil
}
