FROM golang:1.15

LABEL maintainer="Savoiringfaire <savoir@hhra.uk>"

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 2112

ENTRYPOINT ["overseerr-prom-exporter"]