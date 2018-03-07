FROM golang:alpine

RUN apk add --update tzdata \
    bash wget curl git;

# copy is needed here to initially get all deps. Modifying the deps requires rebuilding the image.
COPY . /go/src/github.com/PeterHenell/docker-compose-go-elastic-swagger/api
WORKDIR /go/src/github.com/PeterHenell/docker-compose-go-elastic-swagger/api


#ENV GOPATH=~/go
#ENV GOBIN=$GOPATH/bin
#ENV PATH=$PATH:$GOBIN

RUN mkdir -p $GOPATH/bin && \
    curl https://glide.sh/get | sh && \
    go get github.com/pilu/fresh

RUN glide install
RUN go build ./generated/cmd/myapp-server/

# Glide usage
# https://github.com/Masterminds/glide

EXPOSE 8080

CMD glide update && fresh -c runner.conf ./generated/cmd/myapp-server/main.go

