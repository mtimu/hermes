ADDRESS=
REGISTRY=
TAG=1-0-0
NAME=hermes

emq_up:
	docker run -d --name emqx -p 1883:1883 -p 8081:8081 -p 8083:8083 -p 8883:8883 -p 8084:8084 -p 18083:18083 emqx/emqx:v4.0.0

emq_down:
	docker kill emqx
	
up:
	CGO_ENABLED=0 go build
	docker-compose -f docker-compose.yaml up --build

down:
	docker-compose -f docker-compose.yaml down
	docker-clean


croom:
	curl -X POST http://$(ADDRESS):3000/room \
	-H "Content-Type: application/json" \
	-d '{"host_id": "mehdi"}'

jroom:
	curl -X POST http://$(ADDRESS):3000/room/$(room_id) \
    	-H "Content-Type: application/json" \
    	-d '{"participant_id": "jafar"}'

test:
	curl -X POST http://$(ADDRESS):3000/test

health:
	curl http://$(ADDRESS):3000/healthz

image:
	CGO_ENABLED=0 go build
	docker build -t $(NAME):latest .
	docker tag $(NAME):latest $(REGISTRY)/$(NAME):$(TAG)
	docker push $(REGISTRY)/$(NAME):$(TAG)
	docker rmi $(REGISTRY)/$(NAME):$(TAG)
	docker rmi $(NAME):latest

deploy:
	helm install $(NAME) deployment/$(NAME)  --set image.tag=$(TAG) \
	--set image.repository=$(REGISTRY)/$(NAME)

deployd:
	helm uninstall $(NAME)