# docker-compose-go-elastic-swagger
Skeleton for running Go with elasticsearch generated from swagger spec


# Prerequisites

 * Docker
 * Docker-compose

Tested on Ubuntu 16

# Getting started
Follow these steps when checking out and starting this example.

The service should now be served using fresh (https://github.com/pilu/fresh) making it hot reload on code changes.
```
git clone git@github.com:PeterHenell/docker-compose-go-elastic-swagger.git
cd docker-compose-go-elastic-swagger
make swagger-server
make server-up
```
