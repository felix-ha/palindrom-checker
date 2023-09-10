FROM debian:buster

ARG GO_VERSION=go1.19.5.linux-amd64.tar.gz

RUN apt-get update && \
    apt-get install -y wget vim 

WORKDIR /tmp
RUN wget https://go.dev/dl/$GO_VERSION
RUN tar -C /usr/local -xzf /tmp/$GO_VERSION

ENV GOROOT=/usr/local/go
#ENV GOPATH=/app
ENV PATH=$PATH:$GOROOT/bin

WORKDIR /app

COPY . .
RUN go test -v ./...
RUN go build -o palindromcheck main.go
ENV PATH=$PATH:/app
CMD ["palindromcheck"]

#ENTRYPOINT ["/bin/bash"]
