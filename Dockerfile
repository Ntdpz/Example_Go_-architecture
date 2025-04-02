FROM  golang:1.24-alpine

WORKDIR /app

COPY ./cmd/ ./cmd/ 

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd

CMD ["/app/main"]