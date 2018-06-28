 .PHONY: run-web, stop-web rm-web

PWD := $(shell pwd)
USER := $(shell id -u)
GROUP := $(shell id -g)

run-mq:   
	cd hack/swarm && docker-compose -f docker-compose.yaml -p "micro-$(USER)-mq" up
stop-mq:      
	cd hack/swarm && docker-compose -f docker-compose.yaml -p "micro-$(USER)-mq" stop 
rm-mq:    
	cd hack/swarm && docker-compose -f docker-compose.yaml -p "micro-$(USER)-mq" rm 

run-console:
	cd hack/swarm && docker-compose -f docker-compose-golang.yaml -p "micro-$(USER)-mq-console" up
stop-console:
	cd hack/swarm && docker-compose -f docker-compose-golang.yaml -p "micro-$(USER)-mq-console" stop 
rm-console:
	cd hack/swarm && docker-compose -f docker-compose-golang.yaml -p "micro-$(USER)-mq-console" rm 

build-golang:
	cd hack/dockerfile && docker build -f ./Dockerfile-golang -t mq/golang:1.10 .
