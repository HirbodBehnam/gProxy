package api

import (
	"context"
	"gProxy/pkg/proto"
	"io"
	"net"
	"strconv"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	// The RPC client that can initiate connections to remote server
	RPCClient proto.ProxyServiceClient
	// The listener which accepts local connections
	Listener net.Listener
	// The next connection ID
	nextConnectionID uint32
}

func (c *Client) Serve() error {
	for {
		// Wait for a connection
		conn, err := c.Listener.Accept()
		if err != nil {
			return err
		}
		// Accept it
		connectionID := c.nextConnectionID
		c.nextConnectionID++
		go proxyConnection(conn, c.RPCClient, connectionID)
	}
}

// Close the listener
func (c *Client) Close() error {
	return c.Listener.Close()
}

func proxyConnection(conn net.Conn, client proto.ProxyServiceClient, connectionID uint32) {
	logger := log.WithField("id", connectionID)
	logger.Debug("Accepted connection from ", conn.RemoteAddr())
	defer conn.Close()
	// Create the stream
	md := metadata.New(map[string]string{
		connectionIDMetadata: strconv.FormatUint(uint64(connectionID), 10),
	})
	stream, err := client.Proxy(metadata.NewOutgoingContext(context.Background(), md))
	if err != nil {
		logger.Error("cannot connect to proxy server: ", err)
		return
	}

	// Proxy data
	waitc := make(chan struct{})
	go func() { // remote reader goroutine
		for {
			in, err := stream.Recv()
			if err == io.EOF { // Normal close
				logger.Debug("gRPC read side closed")
				break
			}
			if err != nil {
				logger.WithError(err).Debug("cannot read data from gRPC")
				break
			}
			_, err = conn.Write(in.GetData())
			if err != nil {
				logger.WithError(err).Debug("gRPC read side error")
				break
			}
		}
		close(waitc)
		_ = conn.Close()
	}()

	// Read from socket
	buffer := make([]byte, 32*1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			logger.WithError(err).Debug("socket read error")
			break
		}
		err = stream.Send(&proto.TCPStreamPacket{Data: buffer[:n]})
		if err != nil {
			logger.WithError(err).Debug("gRPC writer error")
			break
		}
	}

	// Done
	_ = stream.CloseSend()
	<-waitc
}
