FROM golang:1.25.5-alpine3.23 as base
RUN apk update
WORKDIR /src/ultimate-tic-tac-toe
COPY go.mod go.sum ./
COPY . .
RUN ls
RUN go build -o ultimate-tic-tac-toe ./cmd/app

FROM alpine:3.20 as binary
RUN apk update && apk add --no-cache tzdata
COPY --from=base /src/ultimate-tic-tac-toe/ultimate-tic-tac-toe .
EXPOSE 8080
CMD ["./ultimate-tic-tac-toe"]