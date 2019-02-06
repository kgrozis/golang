FROM golang:1.8

WORKDIR /go/src/app
COPY . .

ADD src/ /go/src/

#RUN go get -d -v ./...
#RUN go install -v ./...
RUN apt-get update && apt-get install vim -y && rm -rf /var/lib/apt/lists/*

EXPOSE 6001

#CMD ["app"]