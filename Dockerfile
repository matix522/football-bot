FROM golang:1.22.3-alpine as build

WORKDIR /app

COPY . /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o /football-bot

COPY *.go ./

FROM alpine:latest AS football-bot

COPY --from=build /football-bot /football-bot
COPY .env .env

CMD ["/football-bot"]
