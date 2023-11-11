FROM golang:1.20

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

CMD ["go", "cmd/main.go"]

CMD ["tail", "-f", "/dev/null"]