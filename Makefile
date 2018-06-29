 .PHONY: run-web, stop-web rm-web

PWD := $(shell pwd)
USER := $(shell id -u)
GROUP := $(shell id -g)

run-mq:   
	cd hack/swarm && docker-compose -f docker-compose.yaml -p "micro-$(USER)-mq" up -d
stop-mq:      
	cd hack/swarm && docker-compose -f docker-compose.yaml -p "micro-$(USER)-mq" stop 
rm-mq:    
	cd hack/swarm && docker-compose -f docker-compose.yaml -p "micro-$(USER)-mq" rm 

run-kafka: build-golang
	cd hack/swarm && docker-compose -f docker-compose-kafka.yaml -p "micro-$(USER)-mq-kafka" up
stop-kafka:
	cd hack/swarm && docker-compose -f docker-compose-kafka.yaml -p "micro-$(USER)-mq-kafka" stop 
rm-kafka:
	cd hack/swarm && docker-compose -f docker-compose-kafka.yaml -p "micro-$(USER)-mq-kafka" rm 

build-golang:
	cd hack/dockerfile && docker build -f ./Dockerfile-golang -t mq/golang:1.10 .
