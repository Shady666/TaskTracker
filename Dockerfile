FROM golang:1.21 AS builder
LABEL authors="Danil"
WORKDIR /taskTracker
COPY ./ ./
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -trimpath -o taskTracker ./cmd/main.go
CMD /wait && ./taskTracker