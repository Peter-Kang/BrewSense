FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY *.go ./
RUN go build -o gindoc

EXPOSE 8080

CMD ["./gindoc"]