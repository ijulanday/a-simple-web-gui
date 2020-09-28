# cool debian+go docker image
FROM balenalib/beaglebone-green-debian-golang:1-sid AS build

# from dockerhub documentation
WORKDIR /go/src/app

# move our app stuff to container workdir
COPY /app ./

RUN go build

# stuff from example (https://github.com/balena-io-examples/balena-go-hello-world/blob/master/Dockerfile.template)
FROM balenalib/beaglebone-green-debian-golang:stretch

COPY --from=build /go/src/app/ .

CMD ./app