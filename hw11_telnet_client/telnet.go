package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Client struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{address, timeout, in, out, nil}
}

func (c *Client) Connect() (err error) {
	c.conn, err = net.DialTimeout("tcp", c.address, c.timeout)
	return err
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Send() error {
	_, err := io.Copy(c.conn, c.in)
	return err
}

func (c *Client) Receive() error {
	_, err := io.Copy(c.out, c.conn)
	return err
}
