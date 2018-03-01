glide-install: glide-clean
	glide install

glide-clean:
	rm -rf vendor

swagger-server: swagger-clean
	docker-compose up swagger
	sudo chown -R ${USER} api/generated
	sudo chgrp -R ${USER} api/generated

swagger-clean:
	rm -rf api/generated/models api/generated/restapi

docker-nuke:
	docker ps -a -q | xargs docker stop
	docker ps -a -q | xargs docker rm --force
	docker images -q | xargs docker rmi --force

server-up: swagger-server
	docker-compose up api
