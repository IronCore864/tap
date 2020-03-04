FROM golang:alpine AS build-env
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git
WORKDIR $GOPATH/src/github.com/ironcore864/unitet
COPY . .
RUN go get ./... && go build -o unitet

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/ironcore864/unitet/unitet /app/
ENTRYPOINT sh
