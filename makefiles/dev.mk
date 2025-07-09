.PHONY: monitoring-up
monitoring-up:
	@echo "Starting monitoring..."
	@docker-compose -f _harness/monitoring/docker-compose.yaml up -d

.PHONY: monitoring-down
monitoring-down:
	@echo "Stopping monitoring..."
	@docker-compose -f _harness/monitoring/docker-compose.yaml down

.PHONY: stack-up
stack-up:
	@echo "Starting harness..."
	@docker-compose -f _harness/stack/docker-compose.yaml up -d

.PHONY: stack-down
stack-down:
	@echo "Stopping harness..."
	@docker-compose -f _harness/stack/docker-compose.yaml down


.PHONY: run-backend-server
run-backend-server:
	@echo "Running backend server..."
	@go run ./backend/cmd/server/main.go