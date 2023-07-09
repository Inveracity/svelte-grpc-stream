.PHONY: clean
clean:
	rm -rf relay/internal/proto
	rm -rf frontend/src/lib/proto

.PHONY: proto
proto: clean
	buf generate

.PHONY: build
build:
	docker build -t inveracity/chat-relay:latest -f docker/relay.dockerfile .
	docker push inveracity/chat-relay:latest
	docker build -t inveracity/chat-frontend:latest -f docker/frontend.dockerfile .
	docker push inveracity/chat-frontend:latest

.PHONY: deploy
deploy:
	nomad job run nomad/relay.nomad
	nomad job run nomad/frontend.nomad
