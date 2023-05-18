FROM golang:1.18-alpine as builder
RUN mkdir /simple-grpc-app
WORKDIR /simple-grpc-app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on
COPY go.mod /simple-grpc-app/
COPY go.sum /simple-grpc-app/
RUN go mod download
COPY . .
RUN go build -a -installsuffix cgo -o simple-grpc-app .
FROM bullseye-slim
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/simple-grpc-app/simple-grpc-app .
CMD ["/app/simple-grpc-app"]
