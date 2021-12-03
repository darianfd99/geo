FROM golang:latest AS build


WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 go test -cover ./...
RUN CGO_ENABLED=0 go build -o ./main ./cmd/main.go


FROM scratch
COPY --from=build /app/main /app/main
EXPOSE 8080 
ENTRYPOINT [ "app/main" ]