package main

import (
	"fmt"
	"gProxy/api"
	"gProxy/pkg/proto"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify either you want to run this program as client or server as the second argument")
		os.Exit(1)
	}

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	switch os.Args[1] {
	case "client":
		runClient()
	case "server":
		runServer()
	default:
		fmt.Println("Invalid operation mode")
		os.Exit(1)
	}

}

func runClient() {
	if len(os.Args) < 4 {
		fmt.Println("Use the program like this:")
		fmt.Println("\t./gProxy client <LISTEN_ADDRESS> <SERVER_ADDRESS> [tls]")
		os.Exit(1)
	}

	// Connect to gRPC server
	var opts []grpc.DialOption
	if len(os.Args) > 4 && os.Args[4] == "tls" {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(nil))}
	} else {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}
	conn, err := grpc.NewClient(os.Args[3], opts...)
	if err != nil {
		log.Fatalf("fail to dial server: %v", err)
	}
	grpcClient := proto.NewProxyServiceClient(conn)

	// Listen on the socket
	l, err := net.Listen("tcp", os.Args[2])
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}

	// Accept connections
	client := &api.Client{
		RPCClient: grpcClient,
		Listener:  l,
	}
	go func() {
		err := client.Serve()
		if err != nil {
			log.Println("serve error:", err)
		}
	}()

	// Wait for shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Graceful shutdown initiated...")
	_ = client.Close()
	_ = conn.Close()
}

func runServer() {
	if len(os.Args) < 4 {
		fmt.Println("Use the program like this:")
		fmt.Println("\t./gProxy server <LISTEN_ADDRESS> <FORWARD_ADDRESS>")
		os.Exit(1)
	}

	apiData := &api.Server{ForwardAddress: os.Args[3]}
	// Create the gRPC server
	var opts []grpc.ServerOption
	if len(os.Args) > 4 && os.Args[4] == "tls" {
		creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
		if err != nil {
			log.WithError(err).Fatalf("cannot load credentials")
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterProxyServiceServer(grpcServer, apiData)
	go func() {
		if err := grpcServer.Serve(getListener(os.Args[2])); err != nil {
			log.Fatalf("Cannot serve: %s", err)
		}
	}()

	// Wait for shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Graceful shutdown initiated...")
	grpcServer.GracefulStop()
}

// getListener will start a listener based on environment variables
func getListener(address string) net.Listener {
	// Get protocol
	protocol := "tcp"
	if envProtocol := os.Getenv("LISTEN_PROTOCOL"); envProtocol != "" {
		protocol = envProtocol
	}
	// Listen
	listener, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatalf("cannot listen: %s\n", err)
	}
	return listener
}
