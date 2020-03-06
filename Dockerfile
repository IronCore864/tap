FROM golang:alpine AS build-env
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git
WORKDIR $GOPATH/src/github.com/ironcore864/tap
COPY . .
RUN go get ./... && go build -o tap

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/ironcore864/tap/tap /app/
ENTRYPOINT sh
