#FROM golang:1.8
FROM golang:latest

WORKDIR /go/src/app
COPY . .

ADD src/ /go/src/

#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN apt-get update

EXPOSE 6001

#CMD ["app"]