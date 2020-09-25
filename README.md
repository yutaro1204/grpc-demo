# gRPC Demo

## Provisioning Go Modules

To begin with, `go.mod` is needed to start with GoModules feature.  
Provisioning will be done in the container created with DockerCompose.

1. Build DockerCompose with `docker-compose build`.
2. Enter the container with `docker exec -it echo_web_service /bin/bash`.
3. Initialize Go Module configuration with `go mod init example.com/grpc-demo`.
4. Build with `go build` to provision dependencies.
5. `exit` to leave the container.
6. Run DockerCompose againt with `docker-compose up -d`.

## gRPC

About gRPC for unfamilialized persons.

## Starting gRPC

### `proto` file configuration

There is proto file example.

```proto
syntax = "proto3";
service Sample {
  rpc GetSample (GetSampleMessage) returns (ReturnResponse) {}
}
message GetSampleMessage {
  string threshold = 1;
}
message ReturnResponse {
  string item = 1;
}
```

After creation of this file in `/proto`, the following command should be executed.

```shell
protoc proto/sample.proto --go_out=plugins=grpc:pb/
```

Then the directory `/proto` will be created under `/pb` directory, and `sample.pb.go` will be located in `/pb/proto`.

### Connection

gRPC connection will be happend between a client and a server, so they are needed to be provisioned beforehand.

#### Client

Sample source codes for a client is `/client/client.go`.

##### TLS Configuration

The file `server.crt` is created with `openssl`.  
See the next section `Server` and its `TLS Configuration` where the procedure of `openssl` is described.

```golang
// here the server.crt is set fo TLS connection
creds, err := credentials.NewClientTLSFromFile("server.crt", "")
if err != nil {
  log.Fatalf("Failed to certificate: %v", err)
}

// credentials is the argument for grpc.WithTransportCredentials()
conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithTransportCredentials(creds))
if err != nil {
  log.Fatalf("Client Connection Failed: %v", err)
}
```

#### Server

Sample source codes for a server is `/server/server.go`.

##### TLS Configuration

For TLS Configuration, Private Key, Certificate Signing Request and Server Certification are needed and have to be created with `openssl`.  
When `openssl x509` command is executed, SAN(Subject Alternative Name) and IP parameter are needed, so `san.txt` should be set as `-extfile`.  
The IP in `san.txt` is the port where gRPC server is located.

```shell
$ openssl genrsa 2048 > private.key
$ openssl req -new -key private.key > server.csr
  CommonName => localhost
$ openssl x509 -days 367 -req -signkey private.key < server.csr > server.crt -extfile san.txt
```

In golang implementation, source codes are like below.

```golang
listenPort, err := net.Listen("tcp", ":19003")
if err != nil {
  log.Fatalf("Failed to listen: %v", err)
}

// here the server.crt and private key are set fo TLS connection
cred, err := credentials.NewServerTLSFromFile("server.crt", "private.key")
if err != nil {
  log.Fatalf("Failed to certificate: %v", err)
}
server := grpc.NewServer(grpc.Creds(cred))
```

## Microservice Architecture with ServiceMesh

### AWS AppMesh

### Envoy

## gRPCurl

`grpcurl` provides interaction with gRPC server.  
https://github.com/fullstorydev/grpcurl

#### List

```shell
$ grpcurl -plaintext -import-path . -proto proto/sample.proto localhost:19003 list
```
or
```shell
# reflection.Register(server) is needed in server.go
$ grpcurl -plaintext localhost:19003 list
```

To list more in detail.
```shell
$ grpcurl -plaintext localhost:19003 list Sample
```

#### Describe

```shell
$ grpcurl -plaintext localhost:19003 describe Sample.GetSample
```

#### Request

Request with no parameters.
```shell
$ grpcurl -plaintext -d localhost:19003 Sample.GetSample
```

Request with parameters.
```
$ grpcurl -plaintext -d '{"threshold": "default"}' localhost:19003 Sample.GetSample
```

