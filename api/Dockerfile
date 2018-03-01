FROM golang

# Big thanks to https://github.com/McMenemy/GoDoRP
# Most of this file is use as found, modified only to fit the current use case


# if left blank app will run with dev settings
# to build production image run:
# $ docker build ./api --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env


# copy is needed here to initially get all deps. Modifying the deps requires rebuilding the image.
COPY . /go/src/github.com/PeterHenell/docker-compose-go-elastic-swagger
WORKDIR /go/src/github.com/PeterHenell/docker-compose-go-elastic-swagger


ENV export GOPATH=~/go
ENV export GOBIN=$GOPATH/bin
ENV export PATH=$PATH:$GOBIN

RUN mkdir -p $GOPATH/bin
RUN curl https://glide.sh/get | sh

RUN glide install
RUN go build ./generated/cmd/myapp-server/

# Glide usage
# https://github.com/Masterminds/glide


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