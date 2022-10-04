FROM --platform=linux/amd64 golang:1.19

RUN apt update
RUN apt-get install unzip

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-x86_64.zip
RUN unzip protoc-21.5-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/
RUN mv protoc3/include/* /usr/local/include/

RUN mkdir /grpc_plugins
RUN curl https://packages.grpc.io/archive/2022/09/76588969c02873ace926014c148eaa0debd9eae0-637ad889-b141-4ff9-84bb-0ef63cf63330/protoc/grpc-protoc_linux_x64-1.50.0-dev.tar.gz | tar zxvf - -C /grpc_plugins/

WORKDIR /workspaces/Test
RUN go mod init github.com/shima004/GOALED_Server
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN export PATH="$PATH:$(go env GOPATH)/bin"
