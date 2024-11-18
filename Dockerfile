FROM golang:1.23.3-alpine as builder
WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]