package api

import (
	"gProxy/pkg/proto"
	"io"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	proto.UnimplementedProxyServiceServer
	// Where should we forward traffic for each connection
	ForwardAddress string
}

// Proxy will proxy the connection from a client to forward server
func (api *Server) Proxy(stream proto.ProxyService_ProxyServer) error {
	// Get the connection ID from the stream
	connectionID := "<UNKNOWN>"
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		connectionID = md.Get(connectionIDMetadata)[0]
	}
	logger := log.WithField("id", connectionID)
	logger.Trace("Connection established")
	// Connect to address
	conn, err := net.Dial("tcp", api.ForwardAddress)
	if err != nil {
		log.WithError(err).Error("cannot connect to forward address")
		return err
	}
	// Proxy traffic
	// On a separate goroutine read data from the gRPC stream. The reason we do this
	// is that we can close the gRPC stream ONLY by returning from the current function.
	// Thus, the socket read must happen in the current function because we can control
	// its closure from another goroutine.
	// Also, there is no need to wait for this goroutine. It will exit sometime when
	// the connection is closed.
	go func() {
		defer conn.Close()
		for {
			// Read a packet
			in, err := stream.Recv()
			if err == io.EOF {
				logger.Trace("gRPC reader peacefully closed")
				return
			}
			if err != nil {
				logger.WithError(err).Info("gRPC reader closed with error")
				return
			}
			// Send to socket
			_, err = conn.Write(in.GetData())
			if err != nil {
				logger.WithError(err).Info("gRPC reader socket closed")
				return
			}
		}
	}()
	// Read from socket and send to gRPC stream
	buffer := make([]byte, 32*1024)
	for {
		// Read a packet
		n, err := conn.Read(buffer)
		if err == io.EOF {
			logger.Trace("socket reader peacefully closed")
			return nil
		}
		if err != nil {
			logger.WithError(err).Info("socket reader closed with error")
			return nil
		}
		// Send to gRPC
		err = stream.SendMsg(&proto.TCPStreamPacket{Data: buffer[:n]})
		if err != nil {
			logger.WithError(err).Info("socket reader gRPC error")
			return nil
		}
	}
}
