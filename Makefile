devstack:
	docker-compose up -d

run:
	go mod tidy && CONFIG_PATH=config/local.yml air