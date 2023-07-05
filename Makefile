.PHONY: install
install:
	go build -o ~/.docker/cli-plugins/docker-sleep .

.PHONY: run
run: install
	docker sleep
