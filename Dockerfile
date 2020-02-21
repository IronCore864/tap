FROM golang:alpine AS build-env
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git
WORKDIR $GOPATH/src/github.com/ironcore864/tftpl
COPY . .
RUN go get ./... && go build -o tftpl

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/ironcore864/tftpl/tftpl /app/
ENTRYPOINT sh
