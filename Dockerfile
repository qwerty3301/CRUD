FROM golang:latest

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o my_crud ./cmd/main.go

CMD ["./my_crud"]