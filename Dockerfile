FROM golang:alpine

ENV GO111MODULE=on

RUN mkdir /app 
WORKDIR /app 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["/app/main"]
