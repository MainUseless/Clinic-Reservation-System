#! /bin/bash

docker network create front-back-net
docker network create back-rabbitmq-net
docker network create back-mysql-net


# mysql :-
docker build --file ./containerfile/containerfile_mysql -t mysql-image:1.0 .
docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=passwd -e MYSQL_DATABASE=clinic -v mysql-data:/var/lib/mysql \
    --net back-mysql-net mysql-image:1.0


# rabbitmq :-
docker build --file ./containerfile/containerfile_rabbitmq -t rabbitmq-image:1.0 .
docker run -d --name rabbitmq-container --net back-rabbitmq-net rabbitmq-image:1.0


# backend :- 
docker build --file ./containerfile/containerfile_backend -t clinic-backend-image:1.0 .
docker create --name clinic-backend-container -e port=9999 -e jwt_secret="test" \
    -e rabbitmq_url="amqp://guest:guest@rabbitmq-container:5672/" \
    -e mysql_url="root:passwd@tcp(mysql-container:3306)/clinic" clinic-backend-image:1.0
docker network connect front-back-net clinic-backend-container
docker network connect back-rabbitmq-net clinic-backend-container
docker network connect back-mysql-net clinic-backend-container
docker start clinic-backend-container

#TODO
#frontend :-
docker build --file ./containerfile/containerfile_frontend -t clinic-frontend-image:1.0 .

