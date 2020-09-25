FROM golang:1.15

# install go packages and other dependencies
WORKDIR /go
RUN go get -v github.com/labstack/echo/... \
  && go get -v github.com/pilu/fresh \
  && go get -v github.com/golang/protobuf/protoc-gen-go \
  && go get -v google.golang.org/grpc \
  && go get -v github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
  && go get -v github.com/fullstorydev/grpcurl/... \
  && go install github.com/fullstorydev/grpcurl/cmd/grpcurl \
  && apt-get update \
  && apt-get install unzip

# install protoc
COPY ./protoc-3.13.0-linux-x86_64.zip /tmp/protoc/
WORKDIR /tmp/protoc
RUN unzip protoc-3.13.0-linux-x86_64.zip \
  && mv bin/protoc /usr/bin \
  && mv include/google /usr/include
WORKDIR /tmp
RUN rm -rf protoc

# source
COPY ./src /app
WORKDIR /app/example.com/grpc-demo

# visualize http/2 connection
ENV GODEBUG http2debug=2

EXPOSE 8080
