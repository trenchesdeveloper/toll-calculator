obu:
	@go build -o bin/obu obu/main.go
	@echo "Build obu success"
	@./bin/obu

receiver:
	@go build -o bin/receiver ./data_receiver
	@echo "Build receiver success"
	@./bin/receiver

calculator:
	@go build -o bin/calculator ./distance_calculator
	@echo "Build receiver success"
	@./bin/calculator

.PHONY: obu receiver