FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/ramo798/webapp-api
COPY . .

RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/lib/pq

RUN go build main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/ramo798/webapp-api /app

CMD /app/main $PORT