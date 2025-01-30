FROM --platform=linux/arm64 golang:1.22.1 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /bandapi -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch
COPY --from=builder /bandapi /bandapi
EXPOSE 8080
ENTRYPOINT ["/bandapi"]
