CONTAINER_NAME = golang
CONTAINER_TAG = 0.0.4
CONTAINER = $(CONTAINER_NAME):$(CONTAINER_TAG)
HOME_DIR = $(shell pwd)
.PHONY: all
all:
	docker build --no-cache -t $(CONTAINER) .
	docker run -t -d -p 6001:6001 --name $(CONTAINER_NAME) -v $(HOME_DIR)/go:/go/src/app $(CONTAINER)
	docker exec -it $(CONTAINER_NAME) /bin/bash
.PHONY: clean
clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)