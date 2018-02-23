FROM golang

# Big thanks to https://github.com/McMenemy/GoDoRP
# Most of this file is use as found, modified only to fit the current use case


# if left blank app will run with dev settings
# to build production image run:
# $ docker build ./api --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env

# copy is needed here to initially get all deps. Modifying the deps requires rebuilding the image.
COPY . /go/src/github.com/myapp/api
WORKDIR /go/src/github.com/myapp/api

# added vendor services will need to be included here
# RUN go get ./vendor/database

RUN go get ./
RUN go build

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