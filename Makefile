# mongodb config
DATABASE = mongo
VERSION = 4.4
DOCKER_NAME = linemessage
LOCAL_PORT = 2717
MONGO_INITDB_ROOT_USERNAME = linedb_root
MONGO_INITDB_ROOT_PASSWORD = linedb_pwd
DATABASE_PATH = /home/stevetsaoch/go_projects/cinnox-homework/data:/data/db


docker_init:
	docker run -d --name ${DOCKER_NAME} -p ${LOCAL_PORT}:27017 \
	-e MONGO_INITDB_ROOT_USERNAME=${MONGO_INITDB_ROOT_USERNAME} \
	-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_INITDB_ROOT_PASSWORD} \
	-v ${DATABASE_PATH} ${DATABASE}:${VERSION} --auth

docker_start:
	docker start ${DOCKER_NAME}

server_start:
	go run main.go

docker_stop:
	docker stop ${DOCKER_NAME}

linecommand_init:
	(cd ./linecommand && go build && go install)