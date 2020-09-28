# cool debian+go docker image
FROM balenalib/beaglebone-green-debian-golang:1-sid AS build

# from dockerhub documentation
WORKDIR /go/src/app

# move our app stuff to container workdir
COPY /app ./

# get go-app stuff from github
RUN go mod init github.com/ijulanday/a-simple-web-gui.git
RUN go get -u -v github.com/maxence-charriere/go-app/v7

RUN go build

# stuff from example (https://github.com/balena-io-examples/balena-go-hello-world/blob/master/Dockerfile.template)
FROM balenalib/beaglebone-green-debian-golang:stretch

COPY --from=build /go/src/app/ .

CMD ./app