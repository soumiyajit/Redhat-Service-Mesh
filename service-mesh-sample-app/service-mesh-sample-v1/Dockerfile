FROM golang:1.16-alpine

ENV PORT 8080
EXPOSE 8080

WORKDIR /go/src/app
COPY . .
RUN go mod download

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["service-mesh-sample-v1"]
