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


.PHONY: migrations
migrations:
	@echo "flyway..."
	@docker run --rm \
		--network="todo_network" \
		--volume $(PWD)/backend/data/schema:/flyway/backend/data/schema \
		flyway/flyway:7.2.0-alpine \
			-url="jdbc:postgresql://database:5432/todo" \
			-user="todo" -password="todo" \
			-schemas="public" -defaultSchema="public" \
			-locations=filesystem:/flyway/backend/data/schema \
			migrate info

.PHONY: run-backend-server
run-backend-server:
	@echo "Running backend server..."
	@go run ./backend/cmd/server/main.go