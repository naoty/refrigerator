FROM golang:alpine
WORKDIR /go/src/github.com/naoty/refrigerator
COPY . .
RUN go install
ENTRYPOINT ["refrigerator"]
