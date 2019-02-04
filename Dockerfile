FROM golang:1.8

WORKDIR /go/src/app
COPY . .
COPY src/* /go/src/foo/

ADD src/06.2 /go/src/06.2

#RUN go get -d -v ./...
#RUN go install -v ./...
RUN apt-get update && apt-get install vim -y

EXPOSE 6001

#CMD ["app"]