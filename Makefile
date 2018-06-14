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
