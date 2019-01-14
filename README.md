# go-lang Container

go-lang is a simple micro-service for building go-lang applications based on [go-lang](https://hub.docker.com/_/golang) container.

## GETTING STARTED

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### PREQUISITES
The following prerequisites are required for installing Cox-DC-Build micro-service:
1.	***Docker***: [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/) is required to run the node-js micro-service.  Docker creates environment for isolating applicaiton(s); controlled versioning for softwawre tools, modules, and APIs; and for a reproducible environment on differing OSs.
2.	***Git***: [Git]() is used for Source Code Management (SCM).

### INSTALLING

A step by step series of examples that tell you how to get a development env running

The container image can be manually created with the following command:
```
#  docker build --no-cache -t golang:0.0.1 .
```

The container can be started as an long-lived application with the following commands:
```
#  docker run -t -d -p 6001:6001 --name golang -v $(pwd)/go:/go/src/app golang:0.0.1
```

The user can drop into the containers shell with the following command:
```
#  docker exec -it golang /bin/bash
```

Comple and run golang project:
```
docker run -it --rm --name my-running-app golang:0.0.1
```

### CLEAN UP
A step by step series of examples that tell you how to clean up a running env.

The user can manually clean-up the environment with the following commands:
```
bash# docker stop golang
bash# docker rm golang
```

## RUNNING TESTS

Explain how to run the automated tests for this system