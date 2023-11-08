FROM golang:1.19-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/main.go cmd/main.go
COPY pkg pkg
COPY key key
COPY conf conf
WORKDIR /app/cmd
RUN go build -o /main .

FROM alpine:latest AS deploy
WORKDIR /app
RUN apk update
COPY --from=build /main .
CMD ["./main"]