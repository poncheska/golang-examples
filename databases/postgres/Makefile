./PHONY: test
test:
	@echo "it will take 5 seconds"
	docker-compose up  > /dev/null 2>&1 &
	@echo "----postgres database created"
	sleep 5
	go test -coverprofile=c.out ;\
	go tool cover -html=c.out ;\
	docker-compose down  > /dev/null 2>&1
	@echo "----postgres database deleted"