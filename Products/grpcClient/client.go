package grpcClient

import (
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

type Client struct {
	log              hclog.Logger
	Addr             string
	Port             string
	DialOptions      []grpc.DialOption
	ClientConnection *grpc.ClientConn
}

func NewClient(l hclog.Logger, addr string, port string) *Client {
	return &Client{log: l, Addr: addr, Port: port}
}

func (c *Client) DialUp() error {
	c.log.Info("Connecting to grpc server", "Addr", c.Addr, "Port", c.Port)
	client, err := grpc.Dial(c.Addr+":"+c.Port, c.DialOptions...)
	c.ClientConnection = client
	return err
}

//WithDialOption adds dial option to the connection.
//Call to WithDialOption after DialUp won't have any effect
func (c *Client) WithDialOption(do ...grpc.DialOption) {
	for _, opt := range do {
		c.DialOptions = append(c.DialOptions, opt)
	}
}

func (c *Client) Close() {
	c.log.Info("Disconnecting from grpc server")
	c.ClientConnection.Close()
}
