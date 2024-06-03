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
	@echo "Build calculator success"
	@./bin/calculator

aggregator:
	@go build -o bin/aggregator ./aggregator
	@echo "Build aggregator success"
	@./bin/aggregator

.PHONY: obu receiver calculator aggregator