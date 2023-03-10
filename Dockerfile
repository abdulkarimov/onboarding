FROM golang:1.19.0

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest
COPY . .

RUN go get github.com/gofiber/fiber/v2
RUN go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres

RUN go mod tidy