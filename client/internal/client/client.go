package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/arthurshafikov/faraway/client/internal/services"
)

type Client struct {
	conn     net.Conn
	services *services.Services
}

func NewClient(services *services.Services) *Client {
	return &Client{
		services: services,
	}
}

func (c *Client) MakeRequest() {
	if err := c.connectToTheServer(); err != nil {
		log.Fatalln(err)
	}

	message := c.getMessageFromRequest()

	nonce := c.services.ProofOfWork.FindNonce(c.getDataAndDifficultyFromMessage(message))

	fmt.Fprintf(c.conn, "%v\n", nonce)

	quote := c.getMessageFromRequest()
	fmt.Println(quote)

	if err := c.conn.Close(); err != nil {
		log.Fatalln(err)
	}
}

func (c *Client) connectToTheServer() (err error) {
	c.conn, err = net.Dial("tcp", "localhost:8090")

	return
}

func (c *Client) getMessageFromRequest() string {
	msg, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	return msg
}

func (c *Client) getDataAndDifficultyFromMessage(message string) ([]byte, int) {
	args := strings.Split(message, " ")
	difficulty, err := strconv.Atoi(strings.TrimRight(args[1], "\n"))
	if err != nil {
		log.Fatalln(err)
	}

	return []byte(args[0]), difficulty
}
