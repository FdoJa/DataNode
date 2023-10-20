# Variables
IMAGE_NAME = datanode
CONTAINER_NAME = datanode
DOCKERFILE = Dockerfile

# Comandos
build:
	sudo docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

docker-regional:
	sudo docker run -p 8080:80 $(CONTAINER_NAME)
stop:
	sudo docker stop $(CONTAINER_NAME)
rm:
	sudo docker rm $(CONTAINER_NAME)
clean:	sudo stop rm
	sudo docker rmi $(IMAGE_NAME)
logs:
	sudo docker logs -f $(CONTAINER_NAME)
