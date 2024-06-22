FROM golang:1.22.4 as builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN go install github.com/gravityblast/fresh@latest

RUN go build -o awesome .

EXPOSE 8080
CMD ["./awesome"]

