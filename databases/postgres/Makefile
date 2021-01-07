./PHONY: test
test:
	@echo "it will take 5 seconds"
	docker-compose up &
	sleep 5
	@echo "---------------------------------------"
	go run postgres.go
	@echo "---------------------------------------"
	docker-compose down