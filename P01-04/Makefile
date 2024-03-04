.PHONY:
broker:
	@docker compose \
		-f ./compose.yaml \
		up -d broker

.PHONY: test
test: broker
	@echo "Running the tests"
	@go test ./... -coverprofile=coverage_sheet.md
	@docker compose \
		-f ./compose.yaml \
		down broker

.PHONY: run-rabbitmq
run-rabbitmq:
	@docker compose \
		-f ./build/compose.rabbitmq.dev.yaml \
		up --build

.PHONY: run-kafka
run-kafka:
	@docker compose \
		-f ./build/compose.kafka.dev.yaml \
		up --build

.PHONY: coverage
coverage: test
	@go tool cover -html=./utils/coverage_sheet.md

.PHONY: env
env: ./config/.env.tmpl
	cp ./config/.env.tmpl ./config/.env