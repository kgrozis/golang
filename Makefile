CONTAINER_NAME = golang
CONTAINER_TAG = 0.0.2
CONTAINER = $(CONTAINER_NAME):$(CONTAINER_TAG)
HOME_DIR = $(shell pwd)
.PHONY: all
all:
	docker build --no-cache -t $(CONTAINER) .
	docker run -t -d --name $(CONTAINER_NAME) -v $(HOME_DIR)/build:/home/automation/build -v $(HOME_DIR)/data_store:/home/automation/data_store $(CONTAINER)
	docker exec -it $(CONTAINER_NAME) /bin/bash
.PHONY: clean
clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)