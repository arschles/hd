dev:
	docker-compose -p hd up -d postgres
	@echo "sleeping for 5 seconds to wait for the DB to complete startup"
	sleep 5

dev-teardown:
	docker-compose -p hd down
	docker volume prune

docker-build:
	docker build -t arschles/hd .

docker-push:
	docker push arschles/hd
