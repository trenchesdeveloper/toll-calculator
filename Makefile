obu:
	@go build -o bin/obu obu/main.go
	@echo "Build obu success"
	@./bin/obu

receiver:
	@go build -o bin/receiver data_receiver/main.go
	@echo "Build receiver success"
	@./bin/receiver

.PHONY: obu receiver 