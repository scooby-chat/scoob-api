FROM golang:1.19.6

WORKDIR /usr/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o application ./

EXPOSE 9000

CMD ["/usr/app/application"]
