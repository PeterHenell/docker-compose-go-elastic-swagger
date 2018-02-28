FROM golang

# Big thanks to https://github.com/McMenemy/GoDoRP
# Most of this file is use as found, modified only to fit the current use case


# if left blank app will run with dev settings
# to build production image run:
# $ docker build ./api --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env


# copy is needed here to initially get all deps. Modifying the deps requires rebuilding the image.
COPY . /go/src/github.com/docker-compose-go-elastic-swagger/api
WORKDIR /go/src/github.com/docker-compose-go-elastic-swagger/api

# added vendor services will need to be included here
# RUN go get ./vendor/database

ENV export GOPATH=~/go
ENV export GOBIN=$GOPATH/bin
ENV export PATH=$PATH:$GOBIN

RUN mkdir -p $GOPATH/bin
RUN curl https://glide.sh/get | sh

RUN glide install
RUN go build ./cmd/myapp-server

# Glide usage
# https://github.com/Masterminds/glide
#$ glide create                            # Start a new workspace
#$ open glide.yaml                         # and edit away!
#$ glide get github.com/Masterminds/cookoo # Get a package and add to glide.yaml
#$ glide install                           # Install packages and dependencies
## work, work, work
#$ go build                                # Go tools work normally
#$ glide up                                # Update to newest versions of the package

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
    then \
    api; \
    else \
    go get github.com/pilu/fresh && \
    fresh; \
    fi

EXPOSE 8080