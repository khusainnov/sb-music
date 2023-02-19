FROM golang:1.19 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o sb-music main.go

FROM gcr.io/distroless/base-debian11

COPY --from=builder app/sb-music .

EXPOSE 80

CMD ["/sb-music"]
