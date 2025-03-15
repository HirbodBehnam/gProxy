# gProxy

Simple TCP over gRPC tunnel

## Building

You can simply build this app using this command

```bash
go build ./cmd/gProxy
```

## Usage

### Server

You can run the gRPC server without the TLS encryption using the following command:

```bash
./gProxy server <LISTEN_ADDRESS> <FORWARD_ADDRESS>
```

The gRPC server will start on `LISTEN_ADDRESS` and forward traffic to a TCP server at `FORWARD_ADDRESS`.
If you want to start the gRPC server with TLS, you can use the following command:

```bash
./gProxy server <LISTEN_ADDRESS> <FORWARD_ADDRESS> tls
```

Doing so will make the program search for `cert.pem` and `key.pem` files in the current working directory and will use
them for the TLS server.

### Client

Running client is very similar to the server:

```bash
./gProxy server <LISTEN_ADDRESS> <SERVER_ADDRESS>
```

This will create an TCP server on `LISTEN_ADDRESS` and forward the traffic to `SERVER_ADDRESS`. Server will then forward
the traffic to the defined address. You can also tell it to use TLS by adding `tls` to the last argument. Currently,
there is no way to allow insecure connections. 