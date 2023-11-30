#! /bin/bash

docker rm -f mysql-container
docker rm -f rabbitmq-container
docker rm -f clinic-backend-container
docker rm -f clinic-frontend-container

docker network rm -f front-back-net
docker network rm -f back-rabbitmq-net
docker network rm -f back-mysql-net
docker network prune -f

docker rmi -f mysql-image:1.0
docker rmi -f rabbitmq-image:1.0
docker rmi -f clinic-backend-image:1.0
docker rmi -f clinic-frontend-image:1.0