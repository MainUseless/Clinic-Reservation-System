#! /bin/bash

#variables :-
inner_backend_port=9999
exposed_backend_port=9999

inner_frontend_port=8080
exposed_frontend_port=8090

MYSQL_DATABASE=clinic
MYSQL_ROOT_PASSWORD=passwd

jwt_secret="test"


# networks :-
docker network create front-back-net
docker network create back-rabbitmq-net
docker network create back-mysql-net


# mysql :-
docker build --file ./containerfile/containerfile_mysql -t mysql-image:1.0 .
docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD \
    -e MYSQL_DATABASE=$MYSQL_DATABASE -v mysql-data:/var/lib/mysql \
    --net back-mysql-net mysql-image:1.0


# rabbitmq :-
docker build --file ./containerfile/containerfile_rabbitmq -t rabbitmq-image:1.0 .
docker run -d --name rabbitmq-container --net back-rabbitmq-net rabbitmq-image:1.0


# backend :- 
docker build --file ./containerfile/containerfile_backend -t clinic-backend-image:1.0 .
docker create --name clinic-backend-container -p $exposed_backend_port:$inner_backend_port \
    -e port=$inner_backend_port -e jwt_secret=$jwt_secret \
    -e rabbitmq_url="amqp://guest:guest@rabbitmq-container:5672/" \
    -e mysql_url="root:$MYSQL_ROOT_PASSWORD@tcp(mysql-container:3306)/$MYSQL_DATABASE" \
    --net back-mysql-net clinic-backend-image:1.0
docker network connect front-back-net clinic-backend-container
docker network connect back-rabbitmq-net clinic-backend-container
docker start clinic-backend-container


#frontend :-
docker build --file ./containerfile/containerfile_frontend -t clinic-frontend-image:1.0 .
docker run -d --name clinic-frontend-container --net front-back-net -p $exposed_frontend_port:$inner_frontend_port \
    -e REACT_APP_PORT=$inner_frontend_port -e REACT_APP_BACKEND_PORT=$exposed_backend_port -e VITE_BACKEND_URL=http://localhost:$exposed_backend_port \
    clinic-frontend-image:1.0