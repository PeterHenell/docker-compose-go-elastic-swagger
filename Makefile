swagger-server: swagger-clean
	docker-compose up swagger-server
#	sudo chown -R ${USER}:${USER} api/generated
	sudo chown -R ${USER}:${USER} api


swagger-clean:
	sudo rm -rf api/generated/models api/generated/restapi

docker-nuke:
	docker ps -a -q | xargs docker stop
	docker ps -a -q | xargs docker rm --force
	docker images -q | xargs docker rmi --force

server-up: swagger-server
	docker-compose up api

swagger-client:
	docker-compose up swagger-client
	sudo chown -R ${USER}:${USER} web/app/client

web-up: swagger-client
	docker-compose up web

up: swagger-client swagger-server
	docker-compose up
