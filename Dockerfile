FROM golang:1.16-alpine 

WORKDIR /app 

COPY go.mod ./ 
COPY go.sum ./ 

RUN go mod download 

COPY [".", "./"]

EXPOSE 8080 

RUN go build cmd/main.go 

CMD ["./main"]